package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasConfiguracoes = []Rota{
	{
		URI:          "/configuracoes",
		Metodo:       http.MethodGet,
		Funcao:       controllers.CarregarConfiguracoes,
		Autenticacao: true,
	},
}