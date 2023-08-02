package main

import (
	"encoding/json"
	"net/http"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `name`
	}
	json.NewDecoder(r.Body)

	respondWithJSON(w, 200, struct{}{})
}
