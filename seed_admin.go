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
	_, err := userRepository.FindByUsername(username)

	if err == nil {
		return
	}

	password, _ := passwordHash.Hash(os.Getenv("INIT_ADMIN_PASSWORD"))

	_, err = userRepository.Insert(entity.User{
		Username: username,
		Password: password,
	})

	if err != nil {
		fmt.Println(err)
	}
}
