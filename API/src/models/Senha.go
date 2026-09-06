package models

type Senha struct {
	Atual string `json:"atual"`
	Nova  string `json:"nova"`
}
