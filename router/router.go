package router

import (
	"net/http"

	"github.com/brunoOchoa/handlers"
	"github.com/brunoOchoa/prompt"
)

func InitRoutes() {
	http.HandleFunc("/api/prompt", prompt.HandlePrompt)
	http.HandleFunc("/api/users", handlers.ListUsersHandler)

	// http.Handle("/front", http.FileServer(http.Dir("./webui")))
}
