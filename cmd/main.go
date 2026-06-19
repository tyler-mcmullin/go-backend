package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tyler-mcmullin/go-backend/db"
	"github.com/tyler-mcmullin/go-backend/handlers"
)

func main() {

	if err := godotenv.Load("../.env"); err != nil {
		panic(err)
	}

	if err := db.CreateMongoConnection(); err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})

	r.GET("/items/:id", handlers.GetItem)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
