package main

import (
	authController "brodo-demo/api/auth"
	"brodo-demo/api/middleware"
	userController "brodo-demo/api/user"

	"brodo-demo/config"
	"brodo-demo/pkg/postgresql"
	"brodo-demo/pkg/security"
	authService "brodo-demo/service/auth"
	userService "brodo-demo/service/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	conn, err := config.ConnectDB()

	if err != nil {
		panic(err)
	}

	bcryptPasswordHash := security.NewBcryptPasswordHash()
	jwtTokenManager := security.NewJwtTokenManager()

	userRepository := postgresql.NewUserRepositoryPostgreSQL(conn)

	userService := userService.NewUserService(userRepository, bcryptPasswordHash)
	authService := authService.NewAuthService(userRepository, bcryptPasswordHash, jwtTokenManager)

	userController := userController.NewUserContoller(&userService)
	authController := authController.NewAuthController(authService)
	authMiddleware := middleware.NewAuthMiddleware(jwtTokenManager)

	router := gin.New()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/admins", authMiddleware, userController.PostUser)
		v1.POST("/login", authController.PostLogin)
	}

	router.Run(":8000")
}
