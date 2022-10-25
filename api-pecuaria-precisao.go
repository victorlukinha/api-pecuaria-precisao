package main

import (
	"github.com/eucatur/go-toolbox/api"
	"github.com/victorlukinha/api-pecuaria-precisao/animals"
	"github.com/victorlukinha/api-pecuaria-precisao/log"
	"github.com/victorlukinha/api-pecuaria-precisao/owner"
	//"github.com/eucatur/go-toolbox/env"
	//"github.com/eucatur/go-toolbox/redis"
)

func main() {

	api.Make()
	api.Use(log.Middleware())
	api.UseCustomHTTPErrorHandler()

	//adiciona as rotas
	api.ProvideEchoInstance(animals.AddRoutes)
	api.ProvideEchoInstance(owner.AddRoutes)

	//inicia o servidor
	api.Run()

}
