package main

import (
	"github.com/minhluan1590/Youtube_Build_Ecommerce_With_Golang/controllers"
	"github.com/minhluan1590/Youtube_Build_Ecommerce_With_Golang/database"
	"github.com/minhluan1590/Youtube_Build_Ecommerce_With_Golang/middleware"
	"github.com/minhluan1590/Youtube_Build_Ecommerce_With_Golang/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    router := gin.Default()

    // Apply middleware
    router.Use(middleware.CORSMiddleware())
    router.Use(middleware.AuthenticationMiddleware()) // Apply Authentication middleware

    // Initialize routes
    routes.UserRoutes(router)

    // Connect to the database
    database.ConnectDB()

    // Get the port from the environment variable, default to 8080 if not set
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Start the server
    router.Run(":" + port)
}
