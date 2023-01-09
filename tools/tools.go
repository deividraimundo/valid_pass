package tools

import "github.com/deividraimundo/valid_pass/config"

type Tools struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Tools {
	return &Tools{
		cfg: cfg,
	}
}
