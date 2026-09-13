package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"webapp/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"
)

func CarregarConfiguracoes(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	usuarioID, erro := strconv.ParseUint(cookie["id"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroDaAPI{Erro: erro.Error()})
		return
	}

	usuario, erro := buscarUsuarioPorID(r, usuarioID)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroDaAPI{Erro: erro.Error()})
		return
	}

	podeAtualizar := true
	diasRestantes := 0

	if usuario.AtualizadoEm != nil {
		const prazo = 3 * 30 * 24 * time.Hour
		decorrido := time.Since(*usuario.AtualizadoEm)

		if decorrido < prazo {
			podeAtualizar = false
			diasRestantes = int((prazo - decorrido).Hours()/24) + 1
		}
	}

	utils.ExecutarTemplate(w, "configuracoes.html", struct {
		Usuario       models.Usuario
		PodeAtualizar bool
		DiasRestantes int
	}{
		Usuario:       usuario,
		PodeAtualizar: podeAtualizar,
		DiasRestantes: diasRestantes,
	})
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	cookie, _ := cookies.Ler(r)
	usuarioID := cookie["id"]

	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"nick":  r.FormValue("nick"),
		"email": r.FormValue("email"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroDaAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios/%s", config.APIURL, usuarioID)
	response, erro := requisicoes.RequisicaoComAutenticacao(r, http.MethodPut, url, bytes.NewBuffer(usuario))
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

func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	senha, erro := json.Marshal(models.Senha{
		Atual: r.FormValue("senhaAtual"),
		Nova:  r.FormValue("senhaNova"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroDaAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID := cookie["id"]

	url := fmt.Sprintf("%s/usuarios/%s/atualizar-senha", config.APIURL, usuarioID)
	response, erro := requisicoes.RequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(senha))
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

func ExcluirConta(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	usuarioID := cookie["id"]

	url := fmt.Sprintf("%s/usuarios/%s", config.APIURL, usuarioID)
	response, erro := requisicoes.RequisicaoComAutenticacao(r, http.MethodDelete, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroDaAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeErro(w, response)
		return
	}

	cookies.Deletar(w)
	respostas.JSON(w, response.StatusCode, nil)
}