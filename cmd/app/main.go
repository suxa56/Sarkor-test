package main

import (
	"Sarkor-test"
	"Sarkor-test/pkg/handler"
	"Sarkor-test/pkg/repository"
	"Sarkor-test/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(Sarkor_test.Server)
	if err := srv.Run("9090", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
