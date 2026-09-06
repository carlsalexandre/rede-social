package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rota{
	{
		URI:    "/publicacoes",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarPublicacao,
		Autenticacao: true,
	},
	{
		URI:    "/publicacoes",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarPublicacoes,
		Autenticacao: true,
	},
	{
		URI:    "/publicacoes/{publicacaoId}",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarPublicacao,
		Autenticacao: true,
	},
	{
		URI:    "/publicacoes/{publicacaoId}",
		Metodo: http.MethodPut,
		Funcao: controllers.AtualizarPublicacao,
		Autenticacao: true,
	},
	{
		URI:    "/publicacoes/{publicacaoId}",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeletarPublicacao,
		Autenticacao: true,
	},
	{
		URI:    "/usuarios/{usuarioId}/publicacoes",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarPublicacaoDoUsuario,
		Autenticacao: true,
	},
	{
		URI:    "/publicacoes/{publicacaoId}/curtir",
		Metodo: http.MethodPost,
		Funcao: controllers.CurtirPublicacao,
		Autenticacao: true,
	},
	{
		URI:    "/publicacoes/{publicacaoId}/descurtir",
		Metodo: http.MethodPost,
		Funcao: controllers.DescurtirPublicacao,
		Autenticacao: true,
	},
}