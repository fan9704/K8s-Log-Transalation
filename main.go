package main

import (
	"k8s-log-translation/app/config"
	"k8s-log-translation/app/database"
	"k8s-log-translation/app/middleware"
	"k8s-log-translation/app/model"
	"k8s-log-translation/docs"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "k8s-log-translation/docs"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		panic(envErr)
	}

	docs.SwaggerInfo.Title = "K8s Log Translation API"
	docs.SwaggerInfo.Description = "This is a command server on pve vm."
	docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "fan9704.swagger.io"
	// docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	port := os.Getenv("PORT")
	dbConfig := os.Getenv("DB_CONFIG")
	db, ormErr := database.Initialize(dbConfig)
	if ormErr != nil {
		panic(ormErr)
	}
	migrateErr := db.AutoMigrate(&model.User{})
	if migrateErr != nil {
		return 
	}

	app := gin.Default()
	app.GET("/hc", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "health check",
		})
	})
	// Swagger http://127.0.0.1:7780/swagger/index.html#
	url := ginSwagger.URL("http://127.0.0.1:"+port+"/swagger/doc.json")
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	// Cos Config
	app.Use(middleware.CORSMiddleware())
	// Register Router
	config.RouteUsers(app)
	config.RouteCommand(app)
	// Run Server
	err := app.Run(":"+port)
	if err != nil {
		panic(err)
	}
}