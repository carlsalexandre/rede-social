package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasLogin = []Rota{
	{
		URI:          "/",
		Metodo:       http.MethodGet,
		Funcao:       controllers.CarregarTelaLogin,
		Autenticacao: false,
	},
	{
		URI:          "/login",
		Metodo:       http.MethodGet,
		Funcao:       controllers.CarregarTelaLogin,
		Autenticacao: false,
	},
	{
		URI:          "/login",
		Metodo:       http.MethodPost,
		Funcao:       controllers.FazerLogin,
		Autenticacao: false,
	},
	{
		URI:          "/logout",
		Metodo:       http.MethodGet,
		Funcao:       controllers.FazerLogout,
		Autenticacao: false,
	},
}