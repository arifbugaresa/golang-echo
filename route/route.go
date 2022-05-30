package route

import (
	"github.com/labstack/echo/v4"
	midEcho "github.com/labstack/echo/v4/middleware"
	"golang-echo/controller"
	midCustom "golang-echo/middleware"
)

func New() *echo.Echo {

	e := echo.New()
	midCustom.LogMiddleware(e)

	e.GET("/hello", controller.HelloController)
	e.GET("/hello/query-param", controller.HelloControllerWithQueryParam)
	e.POST("/hello/body-form-data", controller.HelloControllerWithBodyFormData)
	e.POST("/hello/body-json", controller.HelloControllerWithBodyJSON)
	e.POST("/hello/body-binding-form-data", controller.HelloControllerWithBodyFormDataBinding)
	e.GET("/hello/:id", controller.HelloControllerWithURLParam)

	e.GET("/users", controller.GetUserController)
	e.POST("/users", controller.PostUserController)

	// With Basic Auth
	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(midEcho.BasicAuth(midCustom.BasicAuthDB))
	eAuthBasic.GET("/users", controller.GetUserController)

	return e
}
