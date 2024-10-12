package main

import (
	"fmt"

	"github.com/dalissongabriel/go-api-rest/database"
	"github.com/dalissongabriel/go-api-rest/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	database.DatabaseConnection()
	routes.HandleRequest()
}
