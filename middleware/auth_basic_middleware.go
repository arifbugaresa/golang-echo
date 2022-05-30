package middleware

import (
	"github.com/labstack/echo/v4"
	"golang-echo/config"
	"golang-echo/model"
)

func BasicAuthDB(email, password string, context echo.Context) (bool, error) {
	var user model.User
	err := config.DB.Where("email = ? AND password = ?", email, password).First(&user).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
