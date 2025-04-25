package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/brunoOchoa/db"
	"github.com/brunoOchoa/models"
)

func CreateUser(nome, status string) error {
	user := models.User{Name: nome, Status: status}
	result := db.DB.Create(&user)
	if result.Error != nil {
		log.Fatal("Erro ao criar usuário:", result.Error)
	}
	log.Println("Usuário criado com sucesso:", user)
	return result.Error
}

// handlers/user.go
func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to /api/users")
	users := []models.User{}
	result := db.DB.Find(&users)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func ListUsers() {
	var users []models.User
	result := db.DB.Find(&users)
	if result.Error != nil {
		fmt.Println("Erro ao buscar usuários:", result.Error)
		return
	}
	for _, u := range users {
		dataCriacao := u.CreatedAt.Format("02/01/2006") // Apenas data
		fmt.Printf("ID: %d | Nome: %s | Status: %s | Criado em: %s\n", u.ID, u.Name, u.Status, dataCriacao)
	}
}

func UpdateUserStatus(id uint, status string) {
	result := db.DB.Model(&models.User{}).Where("id = ?", id).Update("status", status)
	if result.Error != nil {
		fmt.Println("Erro ao atualizar status:", result.Error)
		return
	}
	fmt.Println("Status atualizado com sucesso.")
}

func DeleteUser(id uint) {
	result := db.DB.Delete(&models.User{}, id)
	if result.Error != nil {
		fmt.Println("Erro ao deletar usuário:", result.Error)
		return
	}
	fmt.Println("Usuário deletado com sucesso.")
}
