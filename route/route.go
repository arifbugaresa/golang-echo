package route

import (
	"github.com/labstack/echo/v4"
	"golang-echo/controller"
)

func New() *echo.Echo {

	e := echo.New()

	e.GET("/hello", controller.HelloController)
	e.GET("/hello/query-param", controller.HelloControllerWithQueryParam)
	e.POST("/hello/body-form-data", controller.HelloControllerWithBodyFormData)
	e.POST("/hello/body-json", controller.HelloControllerWithBodyJSON)
	e.POST("/hello/body-binding-form-data", controller.HelloControllerWithBodyFormDataBinding)
	e.GET("/hello/:id", controller.HelloControllerWithURLParam)

	e.GET("/users", controller.GetUserController)
	e.POST("/users", controller.PostUserController)

	return e
}
