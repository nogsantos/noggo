package main

import (
	"fmt"
	"log"

	"github.com/nogsantos/noggo/application/repositories"
	"github.com/nogsantos/noggo/domain"
	"github.com/nogsantos/noggo/framework/utils"
)

func main() {

	db := utils.ConnectDB()

	user := domain.User {
		Name: "Test",
		Email: "test@mail.com",
		Password: "123",
	}

	userRepo := repositories.UserRepositoryDB{Db: db}

	result, insert_error :=userRepo.Insert(&user)

	if insert_error != nil {
		log.Fatalf("Error to insert %v", insert_error)
	}

	fmt.Printf("Created ID %v Name %v Email %v At %v", result.ID, result.Name, result.Email, result.CreatedAt)

}