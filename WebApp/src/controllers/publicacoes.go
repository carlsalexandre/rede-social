package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/config"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
)
 
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
 
	publicacao, erro := json.Marshal(map[string]string{
		"conteudo": r.FormValue("conteudo"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroDaAPI{Erro: erro.Error()})
		return
	}
 
	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	response, erro := requisicoes.RequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(publicacao))
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