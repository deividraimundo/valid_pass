package config

import (
	"flag"
	"fmt"
)

// Config guarda as configurações da API
type Config struct {
	PortService   string
	Authorization string
}

func New() *Config {
	var devMode bool

	flag.BoolVar(&devMode, "devmode", false, "Adicionar esta flag para modo desenvolvimento.")
	flag.Parse()

	if devMode {
		fmt.Println("#### Ambiente de desenvolvimento ####")
	}

	return &Config{
		PortService:   "8080",
		Authorization: "Bearer 3406b8d48918e96a912111467f070e4be22ea2402a1633e14d3ae4febf47598b",
	}
}
