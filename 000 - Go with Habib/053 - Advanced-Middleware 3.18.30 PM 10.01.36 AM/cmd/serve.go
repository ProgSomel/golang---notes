package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve(){
	manager := middleware.NewManager()
	manager.Use(

		middleware.Logger,
		middleware.Preflight,
		middleware.Cors,

	)
	mux := http.NewServeMux()
	wrappedMax := manager.WrapMux(mux);

	initRoutes(mux, manager)
	
	fmt.Println("Server is running on: 3000")

	err := http.ListenAndServe(":3000", wrappedMax) //? return an Error(if there is an Error) or nil(if there is no Error)
	if err != nil{
		fmt.Println("Error starting the server: ", err)
	}
}