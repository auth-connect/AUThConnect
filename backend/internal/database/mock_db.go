package database

import (
	"AUThConnect/internal/models"
	"fmt"
)

var db map[int]models.User

type MockDatabase struct{}

func NewMockDatabase() *MockDatabase {
	db = make(map[int]models.User)

	return &MockDatabase{}
}

func (mdb *MockDatabase) Health() map[string]string {
	return map[string]string{"message": "Hello World"}
}

func (mdb *MockDatabase) GetUsers() ([]models.ReturnUser, error) {
	users := []models.ReturnUser{}
	for i, u := range db {
		user := models.ReturnUser{
			Id:       int64(i),
			Username: u.Username,
			FullName: u.FullName,
			Role:     u.Role,
			Email:    u.Email,
		}
		users = append(users, user)
	}
	return users, nil
}

func (mdb *MockDatabase) GetUser(id int64) (models.ReturnUser, error) {
	u, ok := db[int(id)]
	if !ok {
		return models.ReturnUser{}, fmt.Errorf("not found")
	}
	user := models.ReturnUser{
		Id:       id,
		Username: u.Username,
		FullName: u.FullName,
		Role:     u.Role,
		Email:    u.Email,
	}
	return user, nil
}

func (mdb *MockDatabase) CreateUser(user models.User) (int64, error) {
	id := len(db) + 1
	db[id] = user
	return int64(id), nil
}

func (mdb *MockDatabase) UpdateUser(id int64, user models.User) error {
	return nil
}

func (mdb *MockDatabase) DeleteUser(id int64) error {
	return nil
}

func (mdb *MockDatabase) Close() error {
	db = nil
	return nil
}
