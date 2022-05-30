package controller

import (
	"github.com/labstack/echo/v4"
	"golang-echo/model"
	"net/http"
	"strconv"
)

/*
Memanggil controller
*/
func HelloController(context echo.Context) error {

	user := model.UserExample{
		Name:  "Rudi",
		Email: "rudi@gmail.com",
	}

	return context.JSON(
		http.StatusOK,
		user,
	)
}

/*
Mengirim data dengan url param
*/
func HelloControllerWithURLParam(context echo.Context) error {

	idParam, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}

	user := model.UserExample{
		ID:    idParam,
		Name:  "Rudi",
		Email: "rudi@gmail.com",
	}

	return context.JSON(http.StatusOK, user)
}

func HelloControllerWithQueryParam(context echo.Context) error {

	queryParam := context.QueryParam("search")

	user := model.UserExample{
		Name:  "Rudi",
		Email: "rudi@gmail.com",
	}

	return context.JSON(http.StatusOK, map[string]interface{}{
		"search": queryParam,
		"user":   user,
	})
}

func HelloControllerWithBodyFormData(context echo.Context) error {
	name := context.FormValue("name")
	return context.JSON(http.StatusOK, model.UserExample{
		Name: name,
	})
}

func HelloControllerWithBodyJSON(context echo.Context) error {
	var user model.UserExample

	context.Bind(&user)

	return context.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

func HelloControllerWithBodyFormDataBinding(context echo.Context) error {
	var userBinding model.UserBind

	context.Bind(&userBinding)

	return context.JSON(http.StatusOK, map[string]interface{}{
		"user": userBinding,
	})
}
