package main

import (
	"fmt"
)

// InitCache return cache instance
func InitCache() Cache {
	return make(Cache)
}

// AddUser add user
func (c *Cache) AddUser(login, password, name string, age uint8) error {
	if _, ok := (*c)[login]; ok {
		return fmt.Errorf("User %s is exist, try another login", login)
	}
	(*c)[login] = User{login, password, true, Description{name, age}}
	return nil
}

// DeleteUser delete user
func (c *Cache) DeleteUser(login string) error {
	if _, ok := (*c)[login]; ok {
		delete((*c), login)
		return nil
	}
	return fmt.Errorf("User %s is not exist, please create", login)
}

// LogIn login user
func (c *Cache) LogIn(login, password string) error {
	if user, ok := (*c)[login]; ok {
		if user.IsLogin {
			return fmt.Errorf("You already logged in %s", login)
		}
		if user.Password == password {
			user.IsLogin = true
			(*c)[login] = user
			return nil
		}
		return fmt.Errorf("You typed incorrect password: %s", password)

	}
	return fmt.Errorf("User %s is not exist, please create ", login)
}

// LogOff log off user
func (c *Cache) LogOff(login string) error {
	if user, ok := (*c)[login]; ok {
		if !user.IsLogin {
			return fmt.Errorf("You already logged off %s", login)
		}
		user.IsLogin = false
		(*c)[login] = user
		return nil
	}
	return fmt.Errorf("User %s is not exist, please create ", login)
}
