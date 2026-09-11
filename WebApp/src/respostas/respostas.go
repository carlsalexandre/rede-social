package respostas
 
import (
	"encoding/json"
	"log"
	"net/http"
)
 
type ErroDaAPI struct {
	Erro string `json:"erro`
}
 
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
 
	if statusCode == http.StatusNoContent {
		return
	}
 
	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Println(erro)
	}
}
 
func TratarStatusCodeErro(w http.ResponseWriter, r *http.Response) {
	var erro ErroDaAPI
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}