package main

import (
	"fmt"
	"gin-api/infra/config"
	auth "gin-api/src/auth/routes"
	healthCheck "gin-api/src/healthCheck/routes"
	scopes "gin-api/src/scopes/routes"
	user "gin-api/src/user/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	//infinity := make(chan bool)
	//<-infinity
	// Create a Gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	err := godotenv.Load()
	if err != nil {

		fmt.Println(err)
		//log.Fatal("Error loading .env file")

	}

	config.SetupEnv()
	mongoClient := config.ConnectMongoDB()
	config.Seed(mongoClient)

	healthCheck.SetupHealthCheck(router)
	auth.SetupAuthRoutes(router, mongoClient)
	user.SetupUserRoutes(router, mongoClient)
	scopes.SetupScopeRoutes(router, mongoClient)
	// Start serving the application
	router.Run() // default listens and serves on 0.0.0.0:8080
}
