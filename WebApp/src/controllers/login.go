package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/respostas"
)

func FazerLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroDaAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.APIURL)
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

	var autenticacaoDados models.AutenticacaoDados
	if erro = json.NewDecoder(response.Body).Decode(&autenticacaoDados); erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroDaAPI{Erro: erro.Error()})
		return
	}

	if erro = cookies.Salvar(w, autenticacaoDados.ID, autenticacaoDados.Token); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroDaAPI{Erro: erro.Error()})
		return
	}

	respostas.JSON(w, http.StatusOK, nil)
}
