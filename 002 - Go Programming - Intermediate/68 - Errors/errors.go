package main

import (
	"errors"
	"fmt"
)

type myError struct {
	message string
}

func readConfig() error {
	return errors.New("Config Error")
}

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err)
	}
	return nil
}

func main() {
	if err := readData(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data read Successfully")
}