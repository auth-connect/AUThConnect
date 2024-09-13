package server

import (
	"AUThConnect/internal/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	// r.Use(cors.Default())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},                                                                // Allowed origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                                              // Allowed methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "hx-current-url", "hx-request", "hx-target"}, // Allowed headers
		AllowCredentials: true,
	}))

	r.GET("/", s.helloWorldHandler)
	r.GET("/health", s.healthHandler)
	r.GET("/users", s.getUsers)
	r.GET("/users/:id", s.getUser)
	r.POST("/users", s.createUser)
	r.PUT("/users/:id", s.updateUser)
	r.DELETE("/users/:id", s.deleteUser)

	return r
}

func (s *Server) helloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) getUsers(c *gin.Context) {
	var requestBody map[string]interface{}
	if c.Request.ContentLength > 0 {
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}

		if len(requestBody) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}
	}

	users, err := s.db.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (s *Server) getUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	// TODO: Handle more errors
	user, err := s.db.GetUser(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (s *Server) createUser(c *gin.Context) {
	var body models.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := validate.Struct(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// TODO: Hash password

	user := models.User{
		Username: body.Username,
		Password: body.Password, // set to hashed password
		FullName: body.FullName,
		Email:    body.Email,
		Role:     body.Role,
	}

	id, err := s.db.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
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

func (s *Server) updateUser(c *gin.Context) {
	var body models.User

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := validate.Struct(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// TODO: Hash password

	user := models.User{
		Username: body.Username,
		Password: body.Password, // set to hashed password
		FullName: body.FullName,
		Email:    body.Email,
		Role:     body.Role,
	}

	if err := s.db.UpdateUser(id, user); err != nil {
		// TODO: Handle more errors
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
}

func (s *Server) deleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	if err := s.db.DeleteUser(id); err != nil {
		// TODO: Handle more errors
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
