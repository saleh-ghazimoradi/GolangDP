package SRP

import "database/sql"

type User struct {
	ID       int
	Username string
	Email    string
}

type UserManager struct {
	users []User
}

func (um *UserManager) AddUser(user User) {
	um.users = append(um.users, user)
}

func (um *UserManager) DeleteUser(id int) {
	for i, u := range um.users {
		if u.ID == id {
			um.users = append(um.users[:i], um.users[i+1:]...)
			break
		}
	}
}

// SaveToDatabase violates the SPR principle due adding another responsibility other than user management
func (um *UserManager) SaveToDatabase(users []User) error {
	return nil
}

// To fix the issue we need to apply the storage functionality to another struct or module

// UserStorage is a struct for storage responsibility
type UserStorage struct {
	db *sql.DB
}

func (us *UserStorage) SaveToDatabase(users []User) error {
	return nil
}
