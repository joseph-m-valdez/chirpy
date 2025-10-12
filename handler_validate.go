package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerValidateChirp(w http.ResponseWriter, req *http.Request) {
	type paramaters struct {
		Body 	string `json:"body"`
	}
	type result struct {
		CleanedBody	string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(req.Body)
	params := paramaters{}
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	cleanMsg := noNaughtyWords(params.Body)
	hadBadWord := cleanMsg != params.Body

	if hadBadWord {
		respondWithJSON(w, http.StatusOK, result{
			CleanedBody: cleanMsg,
		})
		return
	}

	respondWithJSON(w, http.StatusOK, result{
		CleanedBody: params.Body,
	})
}

func noNaughtyWords(msg string) string {
	const blocked = "****"
	naughtyWords := map[string]struct{} {
		"kerfuffle": 	{},
 		"sharbert": 	{},
 		"fornax": 		{},
	}
	

	tokens := strings.Fields(msg)
	for i, token := range tokens {
		if _, bad := naughtyWords[strings.ToLower(token)]; bad {
			tokens[i] = blocked
		}
	}
	return strings.Join(tokens, " ") 
}

