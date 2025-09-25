package rest

import (
	"ecommerce/config"
	"ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Start(cnf config.Config){

	manager := middlewares.NewManager()
	manager.Use(

		middlewares.Logger,
		middlewares.Preflight,
		middlewares.Cors,

	)
	mux := http.NewServeMux()
	wrappedMax := manager.WrapMux(mux);

	initRoutes(mux, manager)

	addr := ":"+ strconv.FormatInt(cnf.HttpPort, 10)
	
	fmt.Println("Server is running on: port", addr)

	err := http.ListenAndServe(addr, wrappedMax) //? return an Error(if there is an Error) or nil(if there is no Error)
	if err != nil{
		fmt.Println("Error starting the server: ", err)
		os.Exit(1)
	}
}