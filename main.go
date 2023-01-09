package main

import (
	"log"

	"github.com/deividraimundo/valid_pass/config"
	"github.com/deividraimundo/valid_pass/services"
	"github.com/deividraimundo/valid_pass/tools"
)

func main() {
	cfg := config.New()

	log.Println("Valid Pass - 1.0")

	tools := tools.New(cfg)

	service := services.New(cfg, tools)
	service.Listen()
}
