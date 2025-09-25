# 054 - Into The Actual Project Structure
## .env
```go
VERSION=1.0.0
SERVICE_NAME=ECOMMERCE
HTTP_PORT=3000
```
![env](assets/image.png)
to bind the .env files with process(virtual computer), we need a library.
[go.dev](https://pkg.go.dev/github.com/joho/Godotenv)

install this package -> github.com/joho/Godotenv
```bash
go get github.com/joho/godotenv
```

after installing **go.mod** file will look like this:
```go
module ecommerce
go 1.24.5
require github.com/joho/godotenv v1.5.1 // indirect
```

**-------------------------------------------------------------------------------------------------------------------**

## config.go
```go
package config

import (
	"fmt"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

var configurations Config

type Config struct{
	Version string
	ServiceName string
	HttpPort int64
}

func loadConfig(){
	err := godotenv.Load()
	if err != nil{
		fmt.Println("Failed to load the env variables: ", err)
		os.Exit(1)
	}
	
	version := os.Getenv("VERSION")
	if version == ""{
		fmt.Println("Version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == ""{
		fmt.Println("Service name is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == ""{
		fmt.Println("Http port is required")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil{
		fmt.Println("Port must be a number")
		os.Exit(1)
	}

	configurations = Config{
		Version: version,
		ServiceName: serviceName,
		HttpPort: port,
	}
}

func GetConfig() Config{
	loadConfig()
	return configurations
}
```

## main.go
```go
package main

import (
	"ecommerce/cmd"
	"ecommerce/config"
	"fmt"
)

//?Here net is a package and http is a sub-package under net package

func main(){
	cnf := config.GetConfig()

	fmt.Println(cnf.Version)
	fmt.Println(cnf.ServiceName)
	fmt.Println(cnf.HttpPort)
	cmd.Serve()
}
```
```bash
go run main.go
1.0.0
ECOMMERCE
3000
Server is running on: 3000
```

**----------------------------------------------------------------------------------------------------------------------**

## serve.go
```go
package cmd

import (
	"ecommerce/config"
	"ecommerce/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Serve(){
	cnf := config.GetConfig()

	manager := middleware.NewManager()
	manager.Use(

		middleware.Logger,
		middleware.Preflight,
		middleware.Cors,

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
```

**----------------------------------------------------------------------------------------------------------------------**

**if we update all fields then we will use **PUT** and if we want to update only a single field then we will use **PATCH** **

**----------------------------------------------------------------------------------------------------------------------**

As per Go Documentation, the code 0 indicates success and all the other values (1-255) indicate failure. Exit code 80-99 are reserved for user errors, whereas code 100-119 are for software or system errors. Some code meanings from those available are: 
0 OK
1 NotOK
80 UsageError
81 UnknownSubcommand
82 RequirementNotMet
83 Forbidden 
84 MovedPermanently
100 InternalError
101 Unavailable