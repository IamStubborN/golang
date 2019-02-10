package main

import (
	"fmt"
)

// Cache database
type Cache map[string]User

// InitCache return cache instance
func InitCache() Databaser {
	return &Cache{}
}

// CheckUser add user
func (c *Cache) CheckUser(login string) error {
	if _, ok := (*c)[login]; ok {
		return nil
	}
	return fmt.Errorf("User %s isn't exist", login)
}

// AddUser add user
func (c *Cache) AddUser(user User) error {
	if err := c.CheckUser(user.Login); err != nil {
		(*c)[user.Login] = user
		return nil
	}
	return fmt.Errorf("User %s already exist", user.Login)
}

// DeleteUser delete user
func (c *Cache) DeleteUser(login string) error {
	if err := c.CheckUser(login); err != nil {
		return err
	}
	delete((*c), login)
	return nil
}

// GetUserByLogin delete user
func (c *Cache) GetUserByLogin(login string) (User, error) {
	var user User
	if err := c.CheckUser(login); err != nil {
		return user, err
	}
	return (*c)[login], nil
}

// UpdateUserByLogin delete user
func (c *Cache) UpdateUserByLogin(login string, user User) error {
	if err := c.CheckUser(login); err != nil {
		return err
	}
	(*c)[login] = user
	return nil
}

// LogIn login user
func (c *Cache) LogIn(login, password string) error {
	user, err := c.GetUserByLogin(login)
	if err != nil {
		return err
	}
	if user.IsLogin {
		return fmt.Errorf("You already logged in %s", login)
	}
	if user.Password == password {
		user.IsLogin = true
		if err := c.UpdateUserByLogin(login, user); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("You typed incorrect password: %s", login)
}

// LogOff log off user
func (c *Cache) LogOff(login string) error {
	user, err := c.GetUserByLogin(login)
	if err != nil {
		return err
	}
	if !user.IsLogin {
		return fmt.Errorf("You already logged off %s", login)
	}
	user.IsLogin = false
	if err := c.UpdateUserByLogin(login, user); err != nil {
		return err
	}
	return nil
}
