package main

import (
	"Sarkor-test"
	"Sarkor-test/pkg/handler"
	"Sarkor-test/pkg/repository"
	"Sarkor-test/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewSQLiteDB()
	if err != nil {
		log.Fatalf("Failed to initialize db: %s", err.Error())
	}
	defer repository.CloseDB(db)

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(Sarkor_test.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
