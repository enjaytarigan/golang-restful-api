package main

import (
	"brodo-demo/entity"
	"brodo-demo/repository"
	"brodo-demo/service/security"
	"fmt"
	"os"
)

func SeedAdmin(userRepository repository.UserRepository, passwordHash security.PasswordHash) {
	username := os.Getenv("INIT_ADMIN_USERNAME")
	passowrd := os.Getenv("INIT_ADMIN_PASSWORD")
	_, err := userRepository.FindByUsername(username)

	if err == nil {
		return
	}

	_, err = userRepository.Insert(entity.User{
		Username: username,
		Password: passowrd,
	})

	if err != nil {
		fmt.Println(err)
	}
}
