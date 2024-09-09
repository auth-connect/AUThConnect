package server

import (
	"AUThConnect/internal/database"
	"AUThConnect/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/", s.HelloWorldHandler)

	r.GET("/health", s.healthHandler)

	r.GET("/users", s.getUsers)

	r.POST("/users", s.createUser)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) getUsers(c *gin.Context) {
	var users []models.User

	c.JSON(http.StatusOK, users)
}

func (s *Server) createUser(c *gin.Context) {
	var body models.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// TODO: Hash password

	db := database.New()

	user := models.User{
		Username: body.Username,
		Password: body.Password, // set to hashed password
		FullName: body.FullName,
		Email:    body.Email,
		Role:     body.Role,
	}

	id, err := db.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	return_user := models.ReturnUser{
		Id:       id,
		Username: body.Username,
		FullName: body.FullName,
		Email:    body.Email,
		Role:     body.Role,
	}

	c.JSON(http.StatusCreated, return_user)
}
