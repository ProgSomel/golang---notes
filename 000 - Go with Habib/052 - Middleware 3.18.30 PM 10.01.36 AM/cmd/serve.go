package cmd

import (
	"ecommerce/global_router"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve(){
	mux := http.NewServeMux()
	
	manager := middleware.NewManager()
	manager.Use(middleware.Logger, middleware.Hudai)

	initRoutes(mux, manager)
	
	fmt.Println("Server is running on: 3000")

	globalRouter := global_router.GlobalRouter(mux)
	err := http.ListenAndServe(":3000", globalRouter) //? return an Error(if there is an Error) or nil(if there is no Error)
	if err != nil{
		fmt.Println("Error starting the server: ", err)
	}
}