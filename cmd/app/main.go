package main

import (
	"Sarkor-test"
	"Sarkor-test/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(Sarkor_test.Server)
	if err := srv.Run("9090", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
