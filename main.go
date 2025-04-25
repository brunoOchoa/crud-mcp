package main

import (
	"fmt"
	"net/http"

	"github.com/brunoOchoa/db"
	"github.com/brunoOchoa/models"
	"github.com/brunoOchoa/router"
)

func main() {
	db.InitDB()
	db.DB.AutoMigrate(&models.User{})
	fmt.Println("Banco de dados conectado com sucesso!")

	// handlers.CreateUser("Bruno", "Ativo")
	// handlers.CreateUser("Su", "Inativo")
	// handlers.CreateUser("Lizzie", "Ativo")

	// handlers.ListUsers()

	router.InitRoutes()

	fs := http.FileServer(http.Dir("./webui/dist"))
	http.Handle("/", fs)

	http.ListenAndServe(":8080", nil)

}
