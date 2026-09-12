package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	neturl "net/url"
	"strconv"
	"strings"
	"webapp/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroDaAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios", config.APIURL)
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroDaAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	termo := strings.TrimSpace(r.URL.Query().Get("usuario"))

	usuarios, erro := buscarUsuariosNaAPI(r, termo)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroDaAPI{Erro: erro.Error()})
		return
	}

	publicacoes, erro := buscarPublicacoesNaAPI(r, termo)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroDaAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecutarTemplate(w, "usuarios.html", struct {
		Usuarios    []models.Usuario
		Publicacoes []models.Publicacao
		Termo       string
		UsuarioID   uint64
	}{
		Usuarios:    usuarios,
		Publicacoes: publicacoes,
		Termo:       termo,
		UsuarioID:   usuarioID,
	})
}

func buscarUsuariosNaAPI(r *http.Request, termo string) ([]models.Usuario, error) {
	url := fmt.Sprintf("%s/usuarios?usuario=%s", config.APIURL, neturl.QueryEscape(termo))
	response, erro := requisicoes.RequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		return nil, erro
	}
	defer response.Body.Close()

	var usuarios []models.Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuarios); erro != nil {
		return nil, erro
	}

	return usuarios, nil
}

func buscarPublicacoesNaAPI(r *http.Request, termo string) ([]models.Publicacao, error) {
	if termo == "" {
		return nil, nil
	}

	url := fmt.Sprintf("%s/publicacoes/buscar?termo=%s", config.APIURL, neturl.QueryEscape(termo))
	response, erro := requisicoes.RequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		return nil, erro
	}
	defer response.Body.Close()

	var publicacoes []models.Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		return nil, erro
	}

	return publicacoes, nil
}

type dadosPerfil struct {
	Logado          bool
	Usuario         models.Usuario
	Publicacoes     []models.Publicacao
	UsuarioID       uint64
	EhPerfilProprio bool
}

func VisualizarPerfil(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroDaAPI{Erro: erro.Error()})
		return
	}

	if _, erro := cookies.Ler(r); erro != nil {
		utils.ExecutarTemplate(w, "perfil.html", dadosPerfil{Logado: false})
		return
	}

	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuarioID)
	response, erro := requisicoes.RequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroDaAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeErro(w, response)
		return
	}

	var usuario models.Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuario); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroDaAPI{Erro: erro.Error()})
		return
	}

	urlPublicacoes := fmt.Sprintf("%s/usuarios/%d/publicacoes", config.APIURL, usuarioID)
	responsePublicacoes, erro := requisicoes.RequisicaoComAutenticacao(r, http.MethodGet, urlPublicacoes, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroDaAPI{Erro: erro.Error()})
		return
	}
	defer responsePublicacoes.Body.Close()

	var publicacoes []models.Publicacao
	if erro = json.NewDecoder(responsePublicacoes.Body).Decode(&publicacoes); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroDaAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioLogadoID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecutarTemplate(w, "perfil.html", dadosPerfil{
		Logado:          true,
		Usuario:         usuario,
		Publicacoes:     publicacoes,
		UsuarioID:       usuarioLogadoID,
		EhPerfilProprio: usuarioID == usuarioLogadoID,
	})
}