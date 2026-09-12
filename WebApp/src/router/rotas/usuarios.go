package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasUsuarios = []Rota{
	{
		URI:          "/criar-conta",
		Metodo:       http.MethodGet,
		Funcao:       controllers.CarregarCadastro,
		Autenticacao: false,
	},
	{
		URI:          "/usuarios",
		Metodo:       http.MethodPost,
		Funcao:       controllers.CriarUsuario,
		Autenticacao: false,
	},
	{
		URI:          "/usuarios",
		Metodo:       http.MethodGet,
		Funcao:       controllers.BuscarUsuarios,
		Autenticacao: true,
	},
	{
		URI:          "/usuarios/{usuarioId}",
		Metodo:       http.MethodGet,
		Funcao:       controllers.VisualizarPerfil,
		Autenticacao: false,
	},
}