package main

import (
	authController "brodo-demo/api/auth"
	categoryController "brodo-demo/api/category"
	"brodo-demo/api/middleware"
	productController "brodo-demo/api/product"
	userController "brodo-demo/api/user"

	"brodo-demo/config"
	"brodo-demo/pkg/postgresql"
	"brodo-demo/pkg/security"
	"brodo-demo/pkg/storage"
	authService "brodo-demo/service/auth"
	categoryService "brodo-demo/service/category"
	productService "brodo-demo/service/product"
	userService "brodo-demo/service/user"

	"github.com/gin-contrib/cors"
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
	s3 := config.NewS3Client()
	storage := storage.NewStorage(s3)

	userRepository := postgresql.NewUserRepositoryPostgreSQL(conn)
	categoryRepository := postgresql.NewCategoryRepositoryPostgreSQL(conn)
	productRepository := postgresql.NewProductRepositoryPostgreSQL(conn)

	userService := userService.NewUserService(userRepository, bcryptPasswordHash)
	authService := authService.NewAuthService(userRepository, bcryptPasswordHash, jwtTokenManager)
	categoryService := categoryService.NewCategoryService(categoryRepository)
	productService := productService.NewProductService(productRepository, categoryRepository, storage)

	userController := userController.NewUserContoller(&userService)
	authController := authController.NewAuthController(authService)
	categoryController := categoryController.NewCategoryController(categoryService)
	productController := productController.NewProductController(productService)

	authMiddleware := middleware.NewAuthMiddleware(jwtTokenManager)

	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers, authorization"},
	}))

	v1 := router.Group("/api/v1")

	v1.POST("/login", authController.PostLogin)

	adminsRouter := v1.Group("/admins", authMiddleware)
	{
		adminsRouter.POST("/", userController.PostUser)
		adminsRouter.POST("/categories", categoryController.PostCategory)
		adminsRouter.GET("/categories", categoryController.GetCategories)
		adminsRouter.PUT("/categories/:categoryId", categoryController.PutCategoryById)
		adminsRouter.DELETE("/categories/:categoryId", categoryController.DeleteCategoryById)
		adminsRouter.POST("/products", productController.PostProduct)
		adminsRouter.GET("/products", productController.GetProducts)
	}

	publicRouter := v1.Group("/pub")
	{
		publicRouter.GET("/categories", categoryController.GetCategories)
		publicRouter.GET("/products", productController.GetProducts)
		publicRouter.GET("/products/:productId", productController.GetProductById)
	}

	router.Run(":8000")
}
