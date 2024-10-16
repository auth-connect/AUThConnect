package main

import (
	"AUThConnect/internal/database"
	"AUThConnect/internal/logger"
	"AUThConnect/internal/mail"
	"context"
	"database/sql"
	"flag"
	"os"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn             string
		maxOpenConns    int
		maxIdleConns    int
		maxConnLifetime string
		maxIdleTime     string
	}
	jwt struct {
		secret string
	}
	smtp struct {
		host     string
		port     string
		username string
		password string
		sender   string
	}
	origin string
}

type application struct {
	config config
	logger *logger.Logger
	models database.Models
	wg     sync.WaitGroup
	mail   mail.Mail
}

func main() {
	var cfg config

	origin := os.Getenv("ORIGIN")
	if origin == "" {
		origin = "http://localhost:4200"
	}

	flag.IntVar(&cfg.port, "port", 8000, "Backend server port")
	flag.StringVar(&cfg.env, "env", os.Getenv("ENV"), "Environment (development|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("DSN"), "PostgreSQL DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxConnLifetime, "db-max-conn-life-time", "5m", "PostgreSQL max connection life time")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection life time")
	flag.StringVar(&cfg.jwt.secret, "jwt-secret", os.Getenv("JWT"), "JWT secret")
	flag.StringVar(&cfg.smtp.host, "smtp-host", os.Getenv("SMTP_HOST"), "SMTP host")
	flag.StringVar(&cfg.smtp.port, "smtp-port", os.Getenv("SMTP_PORT"), "SMTP port")
	flag.StringVar(&cfg.smtp.username, "smtp-username", os.Getenv("SMTP_USERNAME"), "SMTP username")
	flag.StringVar(&cfg.smtp.password, "smtp-password", os.Getenv("SMTP_PASSWORD"), "SMTP password")
	flag.StringVar(&cfg.smtp.sender, "smtp-sender", "AUThConnect <no-reply@auth-connect.gr>", "SMTP sender")
	flag.StringVar(&cfg.origin, "origin", origin, "Alloweded Origins")
	flag.Parse()

	logger := logger.New(os.Stdout, logger.LevelInfo)

	db, err := OpenDB(cfg)
	if err != nil {
		logger.PrintFatal(err, nil)
	}

	defer db.Close()
	logger.PrintInfo("database conenction pool established", nil)

	app := &application{
		config: cfg,
		logger: logger,
		models: database.NewModels(db),
		mail:   mail.New(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.sender),
	}

	err = app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}

func OpenDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	duration, err := time.ParseDuration(cfg.db.maxConnLifetime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(duration)

	duration, err = time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
