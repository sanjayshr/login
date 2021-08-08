package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sanjayshr/login/api/controller"
)

func Run() {

	var err error
	err = godotenv.Load()

	if err != nil {
		log.Fatal("Error env %v", err)
	}

	var server = controller.Server{}
	server.Initialize(os.Getenv("DBName"))
	server.Run(":8000")
}
