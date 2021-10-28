package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	api "lbc/fizzbuzz/handlers"
	"lbc/fizzbuzz/internal/utils"
)

func main() {
	var app api.App
	err := app.Init()
	if err != nil {
		log.Fatal("Invalid conf")
	}

	log.Fatal(app.Router.Run(fmt.Sprintf(":%s", utils.Getenv("PORT", "8081"))))
}
