package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
	"gitlab.com/codelittinc/golang-interview-project-ismael-estrada/controller"
	"gitlab.com/codelittinc/golang-interview-project-ismael-estrada/router"
	"gitlab.com/codelittinc/golang-interview-project-ismael-estrada/service/storage"
	"gitlab.com/codelittinc/golang-interview-project-ismael-estrada/usecase"

	"github.com/go-playground/validator/v10"
)

type config struct {
	DB *storage.ConnectionParams `mapstructure:"db" validate:"required"`
}

func loadConfig(env string) *config {
	viper.AddConfigPath("config")
	viper.SetConfigName(env)
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/app")

	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("Couldn't lead configuration %v\n", err)
	}

	conf := config{}
	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Fatalf("Couldn't unmarshal configuration %v\n", err)
	}

	validate := validator.New()
	err = validate.Struct(conf)
	if err != nil {
		log.Fatalf("Invalid configuration loaded %v\n", err)
	}

	fmt.Printf("Got this config:\n%v\n", conf)

	return &conf
}

func getSQLConnection() *storage.DBService {
	conf := loadConfig("local")

	dbc := storage.ConnectToDB(*conf.DB)
	return storage.New(dbc)
}

func setupServer() *http.Server {
	db := getSQLConnection()
	tagUsecase := usecase.NewTag(db)
	tagController := controller.NewTag(tagUsecase)

	r := router.Setup(*tagController)
	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

	return srv
}

func main() {
	fmt.Println("Server started... maybe")
	setupServer()
}
