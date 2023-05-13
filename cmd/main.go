package main

import (
	"btl/api/router"
	"btl/config"
	"log"
)

func main() {
	config, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatal(err)
		return
	}
	r, err := router.NewRouter()
	if err != nil {
		panic(err)
	}
	r.Run(":" + config.Service.Address)
}
