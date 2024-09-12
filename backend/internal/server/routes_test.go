package server

import (
	"AUThConnect/internal/database"
	"AUThConnect/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var mockDB *database.MockDatabase
var server *Server
var router *gin.Engine
var globalID int = 1

func setup() {
	mockDB = database.NewMockDatabase()
	server = &Server{db: mockDB}
	router = gin.New()

	router.GET("/", server.helloWorldHandler)
	router.GET("/health", server.healthHandler)
	router.GET("/users", server.getUsers)
	router.GET("/users/:id", server.getUser)
	router.POST("/users", server.createUser)
	router.PUT("/users/:id", server.updateUser)
	router.DELETE("/users/:id", server.deleteUser)
}

func teardown() {

}

func reset() {
	mockDB.Reset()
	globalID = 1
}

// Takes a json string as input and returnes the string in numerical order based on the id field
func normalizeJSON(t *testing.T, jsonStr string) string {
	var items []map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &items); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Sort by ID (or any other unique key)
	sort.Slice(items, func(i, j int) bool {
		return items[i]["id"].(float64) < items[j]["id"].(float64)
	})

	normalizedJSON, err := json.Marshal(items)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}
	return string(normalizedJSON)
}

func TestMain(m *testing.M) {
	// Initiallize the shared MockDatabase
	setup()
	// Run the tests
	m.Run()
	// Clean up the MockDatabase
	teardown()
}

func TestRegisterRoutes(t *testing.T) {
	mockDB := database.NewMockDatabase()
	// Initialize the server and register routes
	server := &Server{db: mockDB}
	router := server.RegisterRoutes()

	// Define test cases
	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
		expectedBody   string
	}{
		{"Hello World", "GET", "/", http.StatusOK, `{"message":"Hello World"}`},
		{"Get Users", "GET", "/users", http.StatusOK, `[]`},
		{"Get User by ID", "GET", "/users/1", http.StatusNotFound, `{"error":"user not found"}`},
		{"Create User", "POST", "/users", http.StatusBadRequest, `{"error":"invalid request body"}`},
		{"Update User", "PUT", "/users/1", http.StatusBadRequest, `{"error":"invalid request body"}`},
		{"Delete User", "DELETE", "/users/1", http.StatusNotFound, `{"error":"user not found"}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Equal(t, tt.expectedBody, w.Body.String())
		})
	}
}

func TestHelloWorldHandler(t *testing.T) {
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
	expected := `{"message":"Hello World"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestGetUsers(t *testing.T) {
	reset()

	// Define 2 users
	users := []models.User{
		{Username: "testuser1", Password: "testpassword1", FullName: "test user1", Role: "tester", Email: "test1@example.com"},
		{Username: "testuser2", Password: "testpassword2", FullName: "test user2", Role: "tester", Email: "test2@example.com"},
	}

	// Create 2 users
	for _, u := range users {
		createTestUser(t, u)
	}

	tests := []struct {
		name           string
		expectedStatus int
		expectedBody   string
		requestBody    string
	}{
		{"Get all users", http.StatusOK, `[{"id":1,"username":"testuser1","full_name":"test user1","role":"tester","email":"test1@example.com"},{"id":2,"username":"testuser2","full_name":"test user2","role":"tester","email":"test2@example.com"}]`, ""},
		{"Invalid request with body", http.StatusBadRequest, `{"error":"invalid request body"}`, `{"username":"test"}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/users", bytes.NewBufferString(tt.requestBody))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			if tt.expectedStatus == http.StatusOK {
				// Normalize the JSON arrays for comparison
				expectedBodyJSON := normalizeJSON(t, tt.expectedBody)
				actualBodyJSON := normalizeJSON(t, rr.Body.String())
				assert.Equal(t, expectedBodyJSON, actualBodyJSON)
			} else {
				assert.Equal(t, tt.expectedBody, rr.Body.String())
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	reset()
	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"error":"user not found"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestCreateUser(t *testing.T) {
	reset()

	// Define 2 users
	users := []models.User{
		{Username: "testuser1", Password: "testpassword1", FullName: "test user1", Role: "tester", Email: "test1@example.com"},
		{Username: "testuser2", Password: "testpassword2", FullName: "test user2", Role: "tester", Email: "test2@example.com"},
	}

	// Create 2 users
	for _, u := range users {
		createTestUser(t, u)
	}

	// Define test cases
	tests := []struct {
		name           string
		path           string
		expectedStatus int
		expectedBody   string
	}{
		{"Get the first user", "/users/1", http.StatusOK, `{"id":1,"username":"testuser1","full_name":"test user1","role":"tester","email":"test1@example.com"}`},
		{"Get all users", "/users", http.StatusOK, `[{"id":1,"username":"testuser1","full_name":"test user1","role":"tester","email":"test1@example.com"},{"id":2,"username":"testuser2","full_name":"test user2","role":"tester","email":"test2@example.com"}]`},
		{"Get non-existent user", "/users/999", http.StatusNotFound, `{"error":"user not found"}`},
		// {"Invalid request body: missing field", "1", http.StatusBadRequest, `{"error":"invalid request body"}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
			assert.Equal(t, rr.Code, tt.expectedStatus)
			if tt.path == "/users" {
				// Normalize the JSON arrays for comparison
				expectedBodyJSON := normalizeJSON(t, tt.expectedBody)
				actualBodyJSON := normalizeJSON(t, rr.Body.String())
				assert.Equal(t, expectedBodyJSON, actualBodyJSON)
			} else {
				assert.Equal(t, tt.expectedBody, rr.Body.String())
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	reset()

	// Create a user
	createTestUser(t, models.User{Username: "testuser", Password: "testpassword", FullName: "test user", Role: "tester", Email: "test@example.com"})

	tests := []struct {
		name           string
		userID         string
		expectedStatus int
		expectedBody   string
		requestBody    string
	}{
		{"Update existing user", "1", http.StatusOK, `{"message":"user updated successfully"}`, `{"username":"newtestuser","password":"testpassword","full_name":"test user","role":"tester","email":"test@example.com"}`},
		{"Update non-existing user", "999", http.StatusNotFound, `{"error":"user not found"}`, `{"username":"newtestuser","password":"testpassword","full_name":"test user","role":"tester","email":"test@example.com"}`},
		{"Invalid user ID", "abc", http.StatusBadRequest, `{"error":"invalid user ID"}`, `{"username":"newtestuser","password":"testpassword","full_name":"test user","role":"tester","email":"test@example.com"}`},
		{"Invalid request body: missing field", "1", http.StatusBadRequest, `{"error":"invalid request body"}`, `{"username":"newtestuser","password":"testpassword","full_name":"test user","role":"tester"}`},
		{"Invalid request body: type mismatch", "1", http.StatusBadRequest, `{"error":"invalid request body"}`, `{"username":"newtestuser","password":"testpassword","full_name":"test user","role":"tester", "email": 5}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("PUT", "/users/"+tt.userID, bytes.NewBufferString(tt.requestBody))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, tt.expectedBody)
			}

			if body := rr.Body.String(); body != tt.expectedBody {
				t.Errorf("Handler returned unexpected body: got %v want %v", body, tt.expectedBody)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	reset()

	// Define 2 users
	users := []models.User{
		{Username: "testuser1", Password: "testpassword1", FullName: "test user1", Role: "tester", Email: "test1@example.com"},
		{Username: "testuser2", Password: "testpassword2", FullName: "test user2", Role: "tester", Email: "test2@example.com"},
	}

	// Create 2 users
	for _, u := range users {
		createTestUser(t, u)
	}

	// Define test cases
	tests := []struct {
		name           string
		userID         string
		expectedStatus int
		expectedBody   string
	}{
		{"Delete existing user", "1", http.StatusOK, `{"message":"user deleted successfully"}`},
		{"Delete non-existent user", "999", http.StatusNotFound, `{"error":"user not found"}`},
		{"Invalid user ID", "abc", http.StatusBadRequest, `{"error":"invalid user ID"}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("DELETE", "/users/"+tt.userID, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			if body := rr.Body.String(); body != tt.expectedBody {
				t.Errorf("Handler returned unexpected body: got %v want %v", body, tt.expectedBody)
			}
		})
	}
}

func createTestUser(t *testing.T, u models.User) {
	requestBody := fmt.Sprintf(`{"username":"%s","password":"%s","full_name":"%s","role":"%s","email":"%s"}`, u.Username, u.Password, u.FullName, u.Role, u.Email)
	req, err := http.NewRequest("POST", "/users", bytes.NewBufferString(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	user := models.ReturnUser{Id: int64(globalID), Username: u.Username, FullName: u.FullName, Role: u.Role, Email: u.Email}
	expectedUser := user.String()
	if rr.Body.String() != expectedUser {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expectedUser)
	}

	globalID++
}
