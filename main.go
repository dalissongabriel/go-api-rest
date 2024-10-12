package main

import (
	"fmt"

	"github.com/dalissongabriel/go-api-rest/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
