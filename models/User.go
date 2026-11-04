package models

import (
	"github.com/pankajy-dev/go-rest-app/api-test/utils"
	"github.com/pankajy-dev/go-rest-app/db"
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
