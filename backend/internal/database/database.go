package database

import (
	"AUThConnect/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	GetUsers() ([]models.ReturnUser, error)
	GetUser(id int64) (models.ReturnUser, error)
	CreateUser(user models.InputUser) (int64, error)
	UpdateUser(id int64, user models.InputUser) error
	DeleteUser(id int64) error

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
}

type service struct {
	db *sql.DB
}

var gormDB *gorm.DB

var (
	database   = os.Getenv("DB_DATABASE")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USERNAME")
	port       = os.Getenv("DB_PORT")
	host       = os.Getenv("DB_HOST")
	schema     = os.Getenv("DB_SCHEMA")
	dbInstance *service
)

func initializeDatabase() {
	if err := gormDB.AutoMigrate(models.User{}); err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	err := gormDB.Exec(`ALTER TABLE users ADD CONSTRAINT unique_username_email UNIQUE (username, email);`).Error
	if err != nil {
		log.Fatalf("Error adding unique constraint: %v", err)
	}
}

func (s *service) checkTableConstraintExist(schema, tableName, constraintName string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (
    SELECT 1 FROM information_schema.tables WHERE table_schema = $1 AND table_name = $2 AND constraint_name = $3
  );`
	err := s.db.QueryRow(query, schema, tableName, constraintName).Scan(&exists)

	return exists, err
}

func New() Service {

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", username, password, host, port, database, schema)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}

	gormDB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatal(err)
	}

	// Set connection pool parameters
	db.SetMaxOpenConns(25)                 // Maximum number of open connections
	db.SetMaxIdleConns(25)                 // Maximum number of idle connections
	db.SetConnMaxLifetime(5 * time.Minute) // Maximum lifetime of a connection

	dbInstance = &service{
		db: db,
	}

	dbInstance.checkTableConstraintExist(schema, "users", "unique_username_email")
	initializeDatabase()

	return dbInstance
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf(fmt.Sprintf("db down: %v", err)) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", database)
	return s.db.Close()
}

func (s *service) GetUsers() ([]models.ReturnUser, error) {
	users := []models.ReturnUser{}
	query := `SELECT id, username, full_name, role, email FROM users`

	rows, err := s.db.Query(query)
	if err != nil {
		log.Printf("Error retrieving users: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.ReturnUser
		err := rows.Scan(&user.Id, &user.Username, &user.FullName, &user.Role, &user.Email)
		if err != nil {
			fmt.Printf("Error retrieving a single user: %v", err)
			break
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *service) GetUser(id int64) (models.ReturnUser, error) {
	user := models.ReturnUser{}
	query := `SELECT id, username, full_name, role, email FROM users WHERE id = $1`

	err := s.db.QueryRow(query, id).Scan(&user.Id, &user.Username, &user.FullName, &user.Role, &user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user with id %d not found", id)
		}
		log.Printf("Error retrieving user: %v", err)
		return user, err
	}

	return user, nil
}

func (s *service) CreateUser(user models.InputUser) (int64, error) {
	// var id int64
	// query := `INSERT INTO users (username, hashed_password, full_name, role, email)
	// VALUES ($1, $2, $3, $4, $5)
	// RETURNING id`

	// err := s.db.QueryRow(query, user.Username, user.Password, user.FullName, user.Role, user.Email).Scan(&id)

	u := models.User{
		Username:       user.Username,
		HashedPassword: user.Password,
		FullName:       user.FullName,
		Role:           user.Role,
		Email:          user.Email,
		CreatedAt:      time.Now(),
	}

	//TODO: Use concurrency to check if the `username` or the `email` are in use

	// Attempt to create the user
	if err := gormDB.Create(&u).Error; err != nil {
		if isUniqueConstraintError(err) {
			return 0, fmt.Errorf("Username or email already in use")
		}
		// log.Printf("Error creating user: %v", err)
		return 0, fmt.Errorf("Error creating user: %v", err)
	}

	// if err != nil {
	// 	return 0, err
	// }

	return u.ID, nil
}

func (s *service) UpdateUser(id int64, user models.InputUser) error {
	query := `UPDATE users 
            SET username = $1, hashed_password = $2, full_name = $3, role = $4, email = $5
            WHERE id = $6`

	_, err := s.db.Exec(query, user.Username, user.Password, user.FullName, user.Role, user.Email, id)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}

	log.Printf("Got user with username: %s", user.Username)

	return nil
}

func (s *service) DeleteUser(id int64) error {
	query := `DELETE FROM users WHERE id = $1`

	res, err := s.db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with id %d not found", id)
	}

	log.Printf("Deleted user with id: %d", id)
	return nil
}

func isUniqueConstraintError(err error) bool {
	return errors.Is(err, gorm.ErrDuplicatedKey)
}
