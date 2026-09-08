package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var PaginaPrincipal = Rota{
	URI:          "/home",
	Metodo:       http.MethodGet,
	Funcao:       controllers.CarregarPaginaPrincipal,
	Autenticacao: true,
}
