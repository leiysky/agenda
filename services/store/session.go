package store

import (
	"errors"
)

func UpdateLoginState(user UserType) error {
	client, err := GetClient()
	if err != nil {
		return err
	}
	if !isMatchedUser(user) {
		return errors.New("invalid user")
	}
	session := &client.DB.Session
	session.Username = user.Username
	if err := client.Commit(); err != nil {
		return err
	}
	return client.Dump()
}

func IsLoggedIn() bool {
	client, _ := GetClient()
	if client.DB.Session.Username != "" {
		return true
	}
	return false
}

func GetCurrentUser() UserType {
	client, _ := GetClient()
	user := UserType{client.getSession().Username, ""}
	return user
}

func (cl *ClientType) getSession() *SessionType {
	return &cl.DB.Session
}

func isMatchedUser(user UserType) bool {
	client, _ := GetClient()
	users := client.DB.Collection.Users
	for _, one := range users {
		if one.Username == user.Username && one.Password == user.Password {
			return true
		}
	}
	return false
}
