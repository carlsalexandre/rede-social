package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/respostas"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := authentication.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return 
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return 
	}

	var publicacao models.Publicacao
	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return 
	}

	publicacao.AutorID = usuarioID

	if erro = publicacao.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return 
	}

	db, erro := db.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDePublicacoes(db)
	publicacao.ID, erro = repositorio.Criar(publicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}

	respostas.JSON(w, http.StatusCreated, publicacao)
}

func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := authentication.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return 
	}

	db, erro := db.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDePublicacoes(db)
	publicacoes, erro := repositorio.Buscar(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}

	respostas.JSON(w, http.StatusOK, publicacoes)
}

func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return 
	}

	db, erro := db.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDePublicacoes(db)
	publicacao, erro := repositorio.BuscarPorID(publicacaoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}

	respostas.JSON(w, http.StatusOK, publicacao)
}

func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := authentication.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return 
	}

	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return 
	}

	db, erro := db.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDePublicacoes(db)
	publicacaoSalvaNoBanco, erro := repositorio.BuscarPorID(publicacaoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}

	if publicacaoSalvaNoBanco.AutorID != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não pode atualizar uma publicação de outro usuário"))
		return 
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return 
	}

	var publicacao models.Publicacao
	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return 
	}

	if erro = publicacao.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return 
	}

	if erro = repositorio.Atualizar(publicacaoID, publicacao); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := authentication.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return 
	}

	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return 
	}

	db, erro := db.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDePublicacoes(db)
	publicacaoSalvaNoBanco, erro := repositorio.BuscarPorID(publicacaoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}

	if publicacaoSalvaNoBanco.AutorID != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não pode deletar uma publicação de outro usuário"))
		return 
	}

	if erro = repositorio.Deletar(publicacaoID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func BuscarPublicacaoDoUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return 
	}

	db, erro := db.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDePublicacoes(db)
	publicacoes, erro := repositorio.BuscarPorUsuario(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}

	respostas.JSON(w, http.StatusOK, publicacoes)
}

func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return 
	}

	db, erro := db.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDePublicacoes(db)
	if erro = repositorio.Curtir(publicacaoID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func DescurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return 
	}

	db, erro := db.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDePublicacoes(db)
	if erro = repositorio.Descurtir(publicacaoID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}