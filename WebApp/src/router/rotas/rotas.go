package rotas
 
import (
	"net/http"
	"webapp/src/middlewares"
 
	"github.com/gorilla/mux"
)
 
type Rota struct {
	URI          string
	Metodo       string
	Funcao       func(http.ResponseWriter, *http.Request)
	Autenticacao bool
}
 
func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin
	rotas = append(rotas, rotasUsuarios...)
	rotas = append(rotas, rotasPublicacoes...)
	rotas = append(rotas, PaginaPrincipal)
 
	for _, rota := range rotas {
		
		if rota.Autenticacao {
			router.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
		} else {
			router.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}
 
	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", semCacheDeLongaDuracao(fileServer)))
 
	return router
}
 
func semCacheDeLongaDuracao(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache")
		h.ServeHTTP(w, r)
	})
}
