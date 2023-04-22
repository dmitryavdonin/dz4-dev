package main

import (
	"log"
	"net/http"

	"user-service/controllers"
	"user-service/initializers"
	"user-service/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	server *gin.Engine

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController
)

func init() {
	config, err := initializers.LoadFromEnv()
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadFromEnv()
	if err != nil {
		log.Fatal("FAILED! Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/")
	router.GET("/healthcheck", func(ctx *gin.Context) {
		message := "user_service"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	UserRouteController.UserRoute(router)

	log.Fatal(server.Run(":" + config.ServicePort))
}
