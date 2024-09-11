package server

import (
	"AUThConnect/internal/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)
	r.GET("/users", s.getUsers)
	r.GET("/users/:id", s.getUser)
	r.POST("/users", s.createUser)
	r.PUT("/users/:id", s.updateUser)
	r.DELETE("/users/:id", s.deleteUser)

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
	users := []models.ReturnUser{}

	users, err := s.db.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (s *Server) getUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// TODO: Handle more erros
	user, err := s.db.GetUser(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (s *Server) createUser(c *gin.Context) {
	var body models.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, err)
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

func (s *Server) updateUser(c *gin.Context) {
	var body models.User

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, err)
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
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (s *Server) deleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	if err := s.db.DeleteUser(id); err != nil {
		// TODO: Handle more errors
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
