package main

import (
	"os"

	"github.com/gin-gonic/gin"
godotenv "github.com/joho/godotenv"
)


//load env

func main() {
	godotenv.Load(".env")
	router := gin.Default()
	
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{
			"name":    "Amresh",
			"message": "Hii gin-gonic",
		})
	})

	//port
	var PORT string = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}

	router.Run(":" + PORT)

}
