package main

import (
	"fmt"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("Error in initializing Zap logger")
		return
	}
	/*
		zap logger can contain some buffer by the end of the function,
		so it is better to flush any buffer that the logger contains.
		for that zap logger instance gives us an option of logger.sync()
	*/

	defer logger.Sync()
	
	//? Now, start logging
	logger.Info("This is info message.")
	logger.Info("User logged in.", zap.String("username", "John Doe"), zap.String("method", "GET"))
}