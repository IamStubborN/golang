package main


// Databaser interface for work with DB's
type Databaser interface {
	CheckUser(login string) error
	AddUser(user User) error
	DeleteUser(login string) error
	GetUserByLogin(login string) (User, error)
	UpdateUserByLogin(login string, user User) error
	LogIn(login, password string) error
	LogOff(login string) error
}