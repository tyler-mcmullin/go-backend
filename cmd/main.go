package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/tyler-mcmullin/go-backend/db"
	preorderRoutes "github.com/tyler-mcmullin/go-backend/routes"
)

func main() {

	if err := godotenv.Load("../.env"); err != nil {
		panic(err)
	}

	if err := db.CreateMongoConnection(); err != nil {
		panic(err)
	}

	r := gin.Default()

	preorderRoutes.GetPreorderControllers(r)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
