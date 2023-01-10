package main

import (
	"log"

	"github.com/deividraimundo/valid_pass/config"
	"github.com/deividraimundo/valid_pass/services"
	"github.com/deividraimundo/valid_pass/tools"
)

func main() {

	// Todas as funções New funcionam como um objeto que é instanciado e passado para ele os devidos parâmetros
	// As estruturas com o nome do pacote são para guardar as configurações passadas ao objeto, para que seja possível acessar as mesmas configurações pelo código inteiro

	// cfg guarda as configurações que serão passadas por todo o código
	cfg := config.New()

	log.Println("Valid Pass - 1.0")

	// tools guarda ferramentas/funções
	tools := tools.New(cfg)

	// service é responsável por conter os serviços da API
	service := services.New(cfg, tools)
	service.Listen()
}
