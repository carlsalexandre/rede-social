package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
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

	response, erro := http.Post("http://localhost:5000/login", "application/json", bytes.NewBuffer(usuario))
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

	respostas.JSON(w, http.StatusOK, nil)
}
