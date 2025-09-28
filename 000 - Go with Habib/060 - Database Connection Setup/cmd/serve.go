package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
	"fmt"
	"os"
)

func Serve(){
	cnf := config.GetConfig()
	dbClient, err := db.NewConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	userRepo := repo.NewUserRepo()
	productRepo := repo.NewProductRepo()
	
	server := rest.NewServer(
		cnf,
		product.NewHandler(middlewares.NewMiddlewares(cnf), productRepo), 
		user.NewHandler(*cnf, userRepo),
	)
	server.Start()
}