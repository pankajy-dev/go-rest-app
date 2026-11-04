package models

import (
	"github.com/pankajy-dev/go-rest-app/db"
	"github.com/pankajy-dev/go-rest-app/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	query := "INSERT INTO user (email, password) values (?,?)"

	hashedPass, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := db.DB.Exec(query, user.Email, hashedPass)
	if err != nil {
		return err
	}

	user.Id, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return err
}

func (user User) ValidateCredentials() bool {

	query := "SELECT password FROM user WHERE email = ?"

	row := db.DB.QueryRow(query, user.Email)

	var retrievedPass string
	err := row.Scan(&retrievedPass)
	if err != nil {
		return false
	}
	if utils.Authenticated(retrievedPass, user.Password) {
		return true
	}
	return false
}
