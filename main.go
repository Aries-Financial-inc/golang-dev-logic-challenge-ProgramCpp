package main

import (
	"fmt"

	"github.com/aries-financial-inc/options-service/routes"
)

func main() {
	fmt.Println("listening on port 8080...")
	router := routes.SetupRouter()
	router.Run() // listen and serve on 0.0.0.0:8080
}