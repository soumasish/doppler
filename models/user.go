package models

import "time"

type User struct {
	Username string
	Password string
	Databases []*Database
	Created time.Time
	Updated time.Time
}

func getUsers() []*User{
	return nil
}

func registerUser(user User) (User, error){

}
