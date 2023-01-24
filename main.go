package main

import (
	"library-api-rest/src/database"
	"library-api-rest/src/routes"
)

func main() {
	routes.HandleRequest()
	database.Connection()
}
