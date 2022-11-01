package main

import (
	"github.com/adantasdf/api-go-gin/database"
	"github.com/adantasdf/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
