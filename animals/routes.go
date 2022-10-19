package animals

import (
	//	"bitbucket.org/eucatur/api-eupassagem/user"
	//	"github.com/eucatur/go-toolbox/handler"
	//
	"github.com/eucatur/go-toolbox/handler"
	"github.com/labstack/echo"
)

// AddRoutes adiciona as rotas para viagens
func AddRoutes(e *echo.Echo) {
	public := e.Group("animals")

	public.GET("", FindByFilterHandler, handler.MiddlewareBindAndValidate(&Filter{}))

	//auth.GET("", FindByFilterHandler, handler.MiddlewareBindAndValidate(&Filter{}))
	//auth.GET("/grouped", FindByFilterGroupedHandler, handler.MiddlewareBindAndValidate(&Filter{}))
	//auth.POST("/typemap", TypeMapParameterHandler, handler.MiddlewareBindAndValidate(&TypeMapFilter{}))
}
