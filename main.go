package main

import (
	"fmt"

	"github.com/escarls/go-jwt-robby/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	fmt.Println("Hello 2")
}
