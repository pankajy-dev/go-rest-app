package repository

import (
	"github.com/pankajy-dev/go-rest-app/internal/db"
	"github.com/pankajy-dev/go-rest-app/internal/models"
	"github.com/pankajy-dev/go-rest-app/internal/utils"
)

func SaveUser(user *models.User) error {
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

func ValidateCredentials(user *models.User) bool {
	query := "SELECT id, password FROM user WHERE email = ?"

	row := db.DB.QueryRow(query, user.Email)

	var retrievedPass string
	err := row.Scan(&user.Id, &retrievedPass)
	if err != nil {
		return false
	}

	return utils.Authenticated(retrievedPass, user.Password)
}
