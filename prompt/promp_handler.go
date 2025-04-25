package prompt

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/brunoOchoa/handlers"
)

type PromptRequest struct {
	Prompt string `json:"prompt"`
}

func HandlePrompt(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to /api/prompt")
	body, _ := io.ReadAll(r.Body)
	var req PromptRequest
	json.Unmarshal(body, &req)

	result := ProcessPrompt(req.Prompt)
	fmt.Fprint(w, result)
}

func ProcessPrompt(input string) string {
	input = strings.ToLower(input)

	if strings.Contains(input, "crie um usuário chamado") {
		nome := extrairNome(input)
		status := extrairStatus(input)

		if nome != "" {
			err := handlers.CreateUser(nome, status)
			if err != nil {
				return fmt.Sprintf("Erro ao criar usuário: %v", err)
			}
			return fmt.Sprintf("Usuário %s criado com status %s", nome, status)
		}
		return "Não consegui identificar o nome do usuário."
	}
	return "Comando não reconhecido."
}

func extrairNome(input string) string {
	parts := strings.Split(input, "chamado")
	if len(parts) < 2 {
		return ""
	}
	nomeStatus := strings.TrimSpace(parts[1])
	nome := strings.Split(nomeStatus, "com")[0]
	return strings.TrimSpace(nome)
}

func extrairStatus(input string) string {
	if strings.Contains(input, "status inativo") {
		return "inativo"
	}
	return "ativo"
}
