package main

import (
	authController "brodo-demo/api/auth"
	categoryController "brodo-demo/api/category"
	"brodo-demo/api/middleware"
	userController "brodo-demo/api/user"

	"brodo-demo/config"
	"brodo-demo/pkg/postgresql"
	"brodo-demo/pkg/security"
	authService "brodo-demo/service/auth"
	categoryService "brodo-demo/service/category"
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
	categoryRepository := postgresql.NewCategoryRepositoryPostgreSQL(conn)

	userService := userService.NewUserService(userRepository, bcryptPasswordHash)
	authService := authService.NewAuthService(userRepository, bcryptPasswordHash, jwtTokenManager)
	categoryService := categoryService.NewCategoryService(categoryRepository)

	userController := userController.NewUserContoller(&userService)
	authController := authController.NewAuthController(authService)
	categoryController := categoryController.NewCategoryController(categoryService)

	authMiddleware := middleware.NewAuthMiddleware(jwtTokenManager)

	router := gin.New()

	v1 := router.Group("/api/v1")

	v1.POST("/login", authController.PostLogin)

	adminsRouter := v1.Group("/admins", authMiddleware)
	{
		adminsRouter.POST("/", userController.PostUser)
		adminsRouter.POST("/categories", categoryController.PostCategory)
		adminsRouter.GET("/categories", categoryController.GetCategories)
		adminsRouter.PUT("/categories/:categoryId", categoryController.PutCategoryById)
		adminsRouter.DELETE("/categories/:categoryId", categoryController.DeleteCategoryById)
	}

	router.Run(":8000")
}
