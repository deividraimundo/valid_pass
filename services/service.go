package services

import (
	"log"
	"net/http"

	"github.com/deividraimundo/valid_pass/config"
	"github.com/deividraimundo/valid_pass/tools"
	"github.com/go-chi/chi/v5"
)

type Service struct {
	cfg *config.Config
	t   *tools.Tools
}

func New(cfg *config.Config, t *tools.Tools) *Service {
	return &Service{
		cfg: cfg,
		t:   t,
	}
}

// Listen fica olhando as rotas para quando for realizada uma chamada
func (s *Service) Listen() {
	router := chi.NewRouter()

	router.Get("/", routerDefault)

	router.Post("/verify", s.t.Middleware(s.Verify))

	log.Fatal(http.ListenAndServe(":"+s.cfg.PortService, router))
}

// RouterDefault é criada apenas para dar um OK na rota principal para simbolizar o funcionamento da API
func routerDefault(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{ "message": "OK" }`))
}
