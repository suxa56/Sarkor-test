package main

import (
	"Sarkor-test"
	"log"
)

func main() {
	srv := new(Sarkor_test.Server)
	if err := srv.Run("9090"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
