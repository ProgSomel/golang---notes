package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/review"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	cnf *config.Config
	productHandler *product.Handler
	userHandler *user.Handler
	reviewHandler *review.Handler
}

func NewServer(
	cnf *config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler,
	reviewHandler *review.Handler,
) *Server{
	return &Server{
		cnf: cnf,
		productHandler: productHandler,
		userHandler: userHandler,
		reviewHandler: reviewHandler,
	}
}

func (server *Server) Start(){

	manager := middlewares.NewManager()
	manager.Use(

		middlewares.Logger,
		middlewares.Preflight,
		middlewares.Cors,

	)
	mux := http.NewServeMux()
	wrappedMax := manager.WrapMux(mux);

	// initRoutes(mux, manager)
	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.FormatInt(server.cnf.HttpPort, 10)
	
	fmt.Println("Server is running on: port", addr)

	err := http.ListenAndServe(addr, wrappedMax) //? return an Error(if there is an Error) or nil(if there is no Error)
	if err != nil{
		fmt.Println("Error starting the server: ", err)
		os.Exit(1)
	}
}