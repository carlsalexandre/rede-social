package rotas
 
import (
	"net/http"
	"webapp/src/controllers"
)
 
var rotasPublicacoes = []Rota{
	{
		URI:          "/publicacoes",
		Metodo:       http.MethodPost,
		Funcao:       controllers.CriarPublicacao,
		Autenticacao: true,
	},
	{
		URI:          "/publicacoes/{publicacaoId}",
		Metodo:       http.MethodDelete,
		Funcao:       controllers.DeletarPublicacao,
		Autenticacao: true,
	},
	{
		URI:          "/publicacoes/{publicacaoId}/curtir",
		Metodo:       http.MethodPost,
		Funcao:       controllers.CurtirPublicacao,
		Autenticacao: true,
	},
	{
		URI:          "/publicacoes/{publicacaoId}/descurtir",
		Metodo:       http.MethodPost,
		Funcao:       controllers.DescurtirPublicacao,
		Autenticacao: true,
	},
}