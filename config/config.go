package config

// Config guarda as configurações da API
type Config struct {
	PortService   string
	Authorization string
}

func New() *Config {

	return &Config{
		PortService:   "8080",
		Authorization: "valid_pass",
	}
}
