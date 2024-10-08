package database

import ()

// var db map[int]models.InputUser
//
// type MockDatabase struct{}
//
// func NewMockDatabase() *MockDatabase {
// 	db = make(map[int]models.InputUser)
//
// 	return &MockDatabase{}
// }
//
// func (mdb *MockDatabase) Health() map[string]string {
// 	return map[string]string{"message": "Hello World"}
// }
//
// func (mdb *MockDatabase) GetUsers() ([]models.ReturnUser, error) {
// 	users := []models.ReturnUser{}
// 	for i, u := range db {
// 		user := models.ReturnUser{
// 			Id:       int64(i),
// 			Username: u.Username,
// 			FullName: u.FullName,
// 			Role:     u.Role,
// 			Email:    u.Email,
// 		}
// 		users = append(users, user)
// 	}
// 	return users, nil
// }
//
// func (mdb *MockDatabase) GetUser(id int64) (models.ReturnUser, error) {
// 	u, ok := db[int(id)]
// 	if !ok {
// 		return models.ReturnUser{}, fmt.Errorf("not found")
// 	}
// 	user := models.ReturnUser{
// 		Id:       id,
// 		Username: u.Username,
// 		FullName: u.FullName,
// 		Role:     u.Role,
// 		Email:    u.Email,
// 	}
// 	return user, nil
// }
//
// func (mdb *MockDatabase) CreateUser(user models.InputUser) (int64, error) {
// 	id := len(db) + 1
// 	db[id] = user
// 	return int64(id), nil
// }
//
// func (mdb *MockDatabase) UpdateUser(id int64, user models.InputUser) error {
// 	_, ok := db[int(id)]
// 	if !ok {
// 		return fmt.Errorf("not found")
// 	}
// 	db[int(id)] = user
// 	return nil
// }
//
// func (mdb *MockDatabase) DeleteUser(id int64) error {
// 	_, ok := db[int(id)]
// 	if !ok {
// 		return fmt.Errorf("not found")
// 	}
// 	delete(db, int(id))
// 	return nil
// }
//
// func (mdb *MockDatabase) Close() error {
// 	return nil
// }
//
// func (mdb *MockDatabase) Reset() {
// 	for key := range db {
// 		delete(db, key)
// 	}
// }
