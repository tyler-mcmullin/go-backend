package main

import (
	"net/http"
  	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tyler-mcmullin/GoBackend/mongo"
)

func main(){

	if err := godotenv.Load(); err != nil{
		panic(err)
	}

	if err := mongo.CreateMongoConnection(); err != nil{
		panic(err)
	}

	r := gin.Default()

	r.GET("/test", func (c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message" : "test",
		})
	})

	if err := r.Run(":8080"); err != nil{
		panic(err)
	}
}