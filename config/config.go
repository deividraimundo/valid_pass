package config

import "os"

// Config guarda as configurações da API
type Config struct {
	PortService   string
	Authorization string
}

func New() *Config {

	return &Config{
		PortService:   os.Getenv("PORT_SERVICE"),
		Authorization: "Bearer 3406b8d48918e96a912111467f070e4be22ea2402a1633e14d3ae4febf47598b",
	}
}
