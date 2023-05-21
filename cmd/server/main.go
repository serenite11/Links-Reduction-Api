package main

import (
	"github.com/serenite11/Links-Reduction-Api/internal/app/grpc"
	"github.com/serenite11/Links-Reduction-Api/internal/app/rest"
	"github.com/serenite11/Links-Reduction-Api/internal/app/rest/handlers"
	"github.com/serenite11/Links-Reduction-Api/internal/database"
	"github.com/serenite11/Links-Reduction-Api/internal/repository"
	"github.com/serenite11/Links-Reduction-Api/internal/service"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.SetFormatter(new(log.JSONFormatter))
	db, err := database.NewPostgres()
	if err != nil {
		log.Fatalf("%s", err.Error())
		return
	}
	defer db.Close()
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
	restSrv := new(rest.Server)
	grpcSrv := new(grpc.Server)
	go func() {
		if err := restSrv.Run("8080", handler.InitRoutes()); err != nil {
			log.Fatalf("%s", err.Error())
			return
		}
	}()
	go func() {
		if err := grpcSrv.Run("5500", services); err != nil {
			log.Fatalf("%s", err.Error())
			return
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("Api is Shutting Down...")
}
