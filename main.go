package main

import (
	"fmt"
	"musicAPI/src/controllers"
	"musicAPI/src/repository"
	"musicAPI/src/routes"
	"musicAPI/src/service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	user      = "postgres"
	password  = "handoko"
	dbname    = "musics-api"
	port      = 5432
	ssl       = "disable"
	time      = "Asia/Jakarta"
	localhost = "localhost:5000"
)

func main() {
	psqlInfo := fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		user, password, dbname, port, ssl, time)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	repository := repository.Repository{
		DB: db,
	}

	service := service.Service{
		Repository: repository,
	}

	controller := controllers.Controller{
		Service: service,
	}
	start := routes.Server{
		Router:     gin.Default(),
		Controller: &controller,
	}
	start.StartServer()
}
