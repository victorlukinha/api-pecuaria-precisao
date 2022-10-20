package owner

import "github.com/labstack/echo"

// AddRoutes adiciona as rotas para viagens
func AddRoutes(e *echo.Echo) {
	public := e.Group("owner")
	public.GET("", getOwnerHandler)
	public.PUT("", updateOwnerHandler)
	public.POST("", createOwnerHandler)
	public.DELETE("", deleteOwnerHandler)

}
