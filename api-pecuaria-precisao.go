package main

import (
	"github.com/eucatur/go-toolbox/api"
	"v1/animals"
	"v1/log"
	//"github.com/eucatur/go-toolbox/env"
	//"github.com/eucatur/go-toolbox/redis"
)

func main() {

	api.Make()
	api.Use(log.Middleware())
	api.UseCustomHTTPErrorHandler()
	api.ProvideEchoInstance(animals.AddRoutes)
	api.Run()
}
