package main

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

type (
	// CacheDB database
	CacheDB struct{}
	obj     bson.M
	arr     []interface{}
)

var (
	session *mgo.Session
	db      *mgo.Database
	col     *mgo.Collection
	dbhost  = "127.0.0.1"
)

func init() {
	session, err := mgo.Dial(dbhost)
	if err != nil {
		panic(err)
	}
	db = session.DB("database")
	col = db.C("users")

	session.SetMode(mgo.Monotonic, true)
}

//InitCacheDB return mongodb database
func InitCacheDB() Databaser {
	return &CacheDB{}
}

//CheckUser check user from mongodb database
func (c *CacheDB) CheckUser(login string) error {
	var user User
	err := col.Find(obj{"login": login}).One(&user)
	return err
}

//AddUser add user to mongodb database
func (c *CacheDB) AddUser(user User) error {
	if err := c.CheckUser(user.Login); err != nil {
		return col.Insert(user)
	}
	return fmt.Errorf("User %s already exist", user.Login)
}

//DeleteUser delete user from mongodb database
func (c *CacheDB) DeleteUser(login string) error {
	if err := c.CheckUser(login); err != nil {
		return err
	}
	col.Remove(obj{"login": login})
	return nil
}

//GetUserByLogin get user from mongodb database
func (c *CacheDB) GetUserByLogin(login string) (user User, err error) {
	if err = c.CheckUser(login); err != nil {
		return
	}
	col.Find(obj{"login": login}).One(&user)
	return
}

//UpdateUserByLogin update user from mongodb database
func (c *CacheDB) UpdateUserByLogin(login string, user User) error {
	return col.Update(obj{"login": login}, user)
}

// LogIn login user
func (c *CacheDB) LogIn(login, password string) error {
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
	return fmt.Errorf("You typed incorrect password, %s", login)
}

// LogOff log off user
func (c *CacheDB) LogOff(login string) error {
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

// сonnectionCheck reconect on lose connect
func сonnectionCheck() {
	if err := session.Ping(); err != nil {
		fmt.Println("Lost connection to db!")
		session.Refresh()
		if err := session.Ping(); err == nil {
			fmt.Println("Reconnect to db successful.")
		}
	}
}
