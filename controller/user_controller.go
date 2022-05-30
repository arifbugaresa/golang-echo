package controller

import (
	"github.com/labstack/echo/v4"
	"golang-echo/config"
	"golang-echo/model"
	"net/http"
)

/*
Menambahkan data user
*/
func PostUserController(context echo.Context) (err error) {
	var user model.User

	context.Bind(&user)

	err = config.DB.Save(&user).Error
	if err != nil {
		return context.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return context.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil menambahkan data.",
		"data":    user,
	})
}

/*
Menganmbil semua data user
*/
func GetUserController(context echo.Context) (err error) {
	var users []model.User

	err = config.DB.Find(&users).Error
	if err != nil {
		return context.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return context.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})
}
