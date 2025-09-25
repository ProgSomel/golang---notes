package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/review"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
)

func Serve(){
	cnf := config.GetConfig()
	server := rest.NewServer(
		cnf,
		product.NewHandler(middlewares.NewMiddlewares(cnf)), 
		user.NewHandler(), 
		review.NewHandler(),
	)
	server.Start()
}