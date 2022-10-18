package main

import (
	"github.com/eucatur/go-toolbox/api"
	//"github.com/eucatur/go-toolbox/env"
	//"github.com/eucatur/go-toolbox/redis"
)

func main() {

	api.Make()
	api.Use(log.Middleware())
	api.UseCustomHTTPErrorHandler()
	api.ProvideEchoInstance(sectional.AddRoutes)
}
