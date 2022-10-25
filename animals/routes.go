package animals

import (
	//	"bitbucket.org/eucatur/api-eupassagem/user"
	//	"github.com/eucatur/go-toolbox/handler"
	//
	"github.com/eucatur/go-toolbox/handler"
	"github.com/labstack/echo"
)

// AddRoutes adiciona as rotas para os animais.
func AddRoutes(e *echo.Echo) {
	public := e.Group("animals")

	public.GET("", FindByFilterHandler, handler.MiddlewareBindAndValidate(&Filter{}))
	public.POST("", CreateHandler, handler.MiddlewareBindAndValidate(&Animals{}))
	public.PUT("", UpdateHandler, handler.MiddlewareBindAndValidate(&Animals{}))
	public.DELETE("/:animalId", DeleteHandler)

	// Atualizar o peso do animal
	public.PUT("", UpdatePesoHandler, handler.MiddlewareBindAndValidate(&Peso{}))

}
