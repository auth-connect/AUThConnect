package server

import (
	"AUThConnect/internal/database"
	"AUThConnect/internal/models"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var mockDB *database.MockDatabase
var server *Server
var router *gin.Engine

func TestMain(m *testing.M) {
	// Initiallize the shared MockDatabase
	mockDB = database.NewMockDatabase()
	server = &Server{db: mockDB}
	router = gin.New()

	router.GET("/health", server.HelloWorldHandler)
	router.GET("/users", server.getUsers)
	router.GET("/users/:id", server.getUser)
	router.POST("/users", server.createUser)
	router.PUT("/users/:id", server.updateUser)
	router.DELETE("/users/:id", server.deleteUser)

	// Run the tests
	m.Run()
}

func TestHelloWorldHandler(t *testing.T) {
	// s := &Server{}
	// r := gin.New()
	// Create a test HTTP request
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	// Serve the HTTP request
	router.ServeHTTP(rr, req)
	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	// Check the response body
	expected := "{\"message\":\"Hello World\"}"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestGetUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "[]"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestGetUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "{\"error\":\"User not found\"}"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestCreateUser(t *testing.T) {
	requestBody1 := `{"username":"testuser1","password":"testpassword","full_name":"test user1","role":"tester","email":"test1@example.com"}`
	requestBody2 := `{"username":"testuser2","password":"testpassword","full_name":"test user2","role":"tester","email":"test2@example.com"}`
	req1, err := http.NewRequest("POST", "/users", bytes.NewBufferString(requestBody1))
	if err != nil {
		t.Fatal(err)
	}

	req2, err := http.NewRequest("POST", "/users", bytes.NewBufferString(requestBody2))
	if err != nil {
		t.Fatal(err)
	}

	rr1 := httptest.NewRecorder()
	router.ServeHTTP(rr1, req1)
	if status := rr1.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	rr2 := httptest.NewRecorder()
	router.ServeHTTP(rr2, req2)
	if status := rr2.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	user1 := models.ReturnUser{Id: 1, Username: "testuser1", FullName: "test user1", Role: "tester", Email: "test1@example.com"}
	user2 := models.ReturnUser{Id: 2, Username: "testuser2", FullName: "test user2", Role: "tester", Email: "test2@example.com"}
	expectedUser1 := user1.String()
	if rr1.Body.String() != expectedUser1 {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr1.Body.String(), expectedUser1)
	}
	expectedUser2 := user2.String()
	if rr2.Body.String() != expectedUser2 {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr2.Body.String(), expectedUser2)
	}

	// Test GetUser method after creating user 1
	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code; got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expectedUser1 {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expectedUser1)
	}

	// Test GetUser method after creating user 2
	req, err = http.NewRequest("GET", "/users/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code; got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expectedUser2 {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expectedUser2)
	}

	// Test GetUsers method after creating two users
	req, err = http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned unexpected code; got %v want %v", status, http.StatusOK)
	}

	expectedUsers := "[" + expectedUser1 + "," + expectedUser2 + "]"
	t.Log(rr.Body)
	if rr.Body.String() != expectedUsers {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expectedUsers)
	}
}

func TestUpdateUser(t *testing.T) {
	requestBody := `{"username":"testuser","password":"testpassword","full_name":"test user","role":"tester","email":"test@example.com"}`
	req, err := http.NewRequest("POST", "/users", bytes.NewBufferString(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	user := models.ReturnUser{Id: 3, Username: "testuser", FullName: "test user", Role: "tester", Email: "test@example.com"}
	expectedUser := user.String()
	if rr.Body.String() != expectedUser {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expectedUser)
	}

	requestBody = `{"username":"newtestuser","password":"testpassword","full_name":"test user","role":"tester","email":"test@example.com"}`
	req, err = http.NewRequest("PUT", "/users/3", bytes.NewBufferString(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "{\"message\":\"User updated successfully\"}"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	req, err = http.NewRequest("GET", "/users/3", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code; got %v want %v", status, http.StatusOK)
	}

	expected = `{"id":3,"username":"newtestuser","full_name":"test user","role":"tester","email":"test@example.com"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
