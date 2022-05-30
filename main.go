package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserBind struct {
	ID int `json:"id" form:"id"`
}

func main() {
	e := echo.New()

	e.GET("/hello", HelloController)
	e.GET("/hello/query-param", HelloControllerWithQueryParam)
	e.POST("/hello/body-form-data", HelloControllerWithBodyFormData)
	e.POST("/hello/body-json", HelloControllerWithBodyJSON)
	e.POST("/hello/body-binding-form-data", HelloControllerWithBodyFormDataBinding)
	e.GET("/hello/:id", HelloControllerWithURLParam)

	e.Start(":8080")
}

/*
	Memanggil controller
*/
func HelloController(context echo.Context) error {

	user := User{
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

	user := User{
		ID:    idParam,
		Name:  "Rudi",
		Email: "rudi@gmail.com",
	}

	return context.JSON(http.StatusOK, user)
}

func HelloControllerWithQueryParam(context echo.Context) error {

	queryParam := context.QueryParam("search")

	user := User{
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
	return context.JSON(http.StatusOK, User{
		Name: name,
	})
}

func HelloControllerWithBodyJSON(context echo.Context) error {
	var user User

	context.Bind(&user)

	return context.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

func HelloControllerWithBodyFormDataBinding(context echo.Context) error {
	var userBinding UserBind

	context.Bind(&userBinding)

	return context.JSON(http.StatusOK, map[string]interface{}{
		"user" : userBinding,
	})
}
