package main

import (
	"github.com/joho/godotenv"
	"github.com/serenite11/Links-Reduction-Api/internal/database"
	"github.com/serenite11/Links-Reduction-Api/internal/handlers"
	"github.com/serenite11/Links-Reduction-Api/internal/repository"
	"github.com/serenite11/Links-Reduction-Api/internal/service"
	"github.com/serenite11/Links-Reduction-Api/pkg"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	log.SetFormatter(new(log.JSONFormatter))
	if err := godotenv.Load(); err != nil {
		log.Fatalf("env doesn`t exist:%s", err.Error())
		return
	}
	db, err := database.NewPostgres(database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})
	if err != nil {
		log.Fatalf("%s", err.Error())
		return
	}
	links := database.NewLinkInMemory()
	repo := repository.NewRepository(db)
	services := service.NewService(repo, links)
	handler := handlers.NewHandler(services)
	srv := new(pkg.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("%s", err.Error())
		return
	}
	if err := db.Close(); err != nil {
		log.Fatalf("error db close connection:%s", err.Error())
		return
	}
}
