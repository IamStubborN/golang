package main

import "fmt"

// Cache database of users
type Cache map[string]User

// InitCache return cache instance
func InitCache() Cache {
	return make(Cache)
}

// AddUser add user
func (c Cache) AddUser(login, password string) error {
	if _, ok := c[login]; ok {
		return fmt.Errorf("User %s is exist, try another login", login)
	}
	c[login] = User{login, password, false}
	return nil
}

// DeleteUser delete user
func (c Cache) DeleteUser(login string) error {
	if _, ok := c[login]; ok {
		delete(c, login)
		return nil
	}
	return fmt.Errorf("User %s is not exist, please create", login)
}

// Login login user
func (c Cache) Login(login, password string) error {
	if user, ok := c[login]; ok {
		if user.isLogin {
			return fmt.Errorf("You already logged in %s", login)
		}
		if user.password == password {
			user.isLogin = true
			return nil
		}
		return fmt.Errorf("You typed incorrect password: %s", password)

	}
	return fmt.Errorf("User %s is not exist, please create ", login)
}

//Logoff log off user
func (c Cache) Logoff(login string) error {
	fmt.Println(1)
	if user, ok := c[login]; ok {
		if !user.isLogin {
			return fmt.Errorf("You already logged off %s", login)
		}
		user.isLogin = false
		return nil
	}
	return fmt.Errorf("User %s is not exist, please create ", login)
}
