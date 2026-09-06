package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
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
		Funcao:       controllers.BuscandoUsuario,
		Autenticacao: true,
	},
	{
		URI:          "/usuarios/{usuarioId}",
		Metodo:       http.MethodPut,
		Funcao:       controllers.AtualizarUsuario,
		Autenticacao: true,
	},
	{
		URI:          "/usuarios/{usuarioId}",
		Metodo:       http.MethodDelete,
		Funcao:       controllers.DeleteUsuario,
		Autenticacao: true,
	},
	{
		URI:          "/usuarios/{usuarioId}/seguir",
		Metodo:       http.MethodPost,
		Funcao:       controllers.SeguirUsuario,
		Autenticacao: true,
	},
	{
		URI:          "/usuarios/{usuarioId}/parar-de-seguir",
		Metodo:       http.MethodPost,
		Funcao:       controllers.PararDeSeguirUsuario,
		Autenticacao: true,
	},
	{
		URI:          "/usuarios/{usuarioId}/seguidores",
		Metodo:       http.MethodGet,
		Funcao:       controllers.BuscarSeguidores,
		Autenticacao: true,
	},
	{
		URI:          "/usuarios/{usuarioId}/seguindo",
		Metodo:       http.MethodGet,
		Funcao:       controllers.BuscarSeguindo,
		Autenticacao: true,
	},
	{
		URI:          "/usuarios/{usuarioId}/atualizar-senha",
		Metodo:       http.MethodPost,
		Funcao:       controllers.AtualizarSenha,
		Autenticacao: true,
	},
}
