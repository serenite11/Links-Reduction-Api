package main

import (
	"github.com/joho/godotenv"
	"github.com/serenite11/Links-Reduction-Api/internal/app/http"
	"github.com/serenite11/Links-Reduction-Api/internal/database"
	"github.com/serenite11/Links-Reduction-Api/internal/handlers"
	"github.com/serenite11/Links-Reduction-Api/internal/repository"
	"github.com/serenite11/Links-Reduction-Api/internal/service"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	log.SetFormatter(new(log.JSONFormatter))
	if err := godotenv.Load(); err != nil {
		log.Fatalf("env doesn`t exist:%s", err.Error())
		return
	}
	db, err := database.NewPostgres()
	if err != nil {
		log.Fatalf("%s", err.Error())
		return
	}
	store := os.Getenv("STORE")
	repo := new(repository.Repository)
	if store == "POSTGRES" {
		repo = repository.NewRepositoryPostgres(db)
	} else if store == "IN-MEMORY" {
		repo = repository.NewRepositoryInMemory(map[string]string{})
	} else {
		log.Fatalf("there is no suitable name for the store")
	}
	services := service.NewService(repo)
	handler := handlers.NewHandler(services)
	srv := new(http.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("%s", err.Error())
		return
	}
	if err := db.Close(); err != nil {
		log.Fatalf("error db close connection:%s", err.Error())
		return
	}
}
