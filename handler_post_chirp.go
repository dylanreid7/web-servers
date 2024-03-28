package main

import (
	"github.com/dylanreid7/web-servers/internal/database"
)

func (cfg *apiConfig) handlerPostChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
		Id int `json:"id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}
	cleaned, err := validateChirp(params.Body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	chirp, err := cfg.DB.CreateChirp(cleaned)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create chirp")
		return
	}

	respondWithJSON(w, http.StatusCreated, Chirp{
		ID:   chirp.ID,
		Body: chirp.Body,
	})
}

func validateChirp(body string) {
	if len(params.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}

	badWords := []string{"kerfuffle", "sharbert", "fornax"}
	cleanedBody := cleanText(body, badWords)
	fmt.Println("cleaned: ", cleanedBody)
	respondWithJSON(w, http.StatusCreated, returnVals{
		CleanedBody: cleanedBody,
		Id: id,
	})
	database.CreateChirp(cleanedBody)
	id++
}

func cleanText(text string, badWords []string) string {
	words := strings.Split(text, " ")
	for i := 0; i < len(words); i++ {
		lower := strings.ToLower(words[i])
		for n := range badWords {
			if lower == badWords[n] {
				words[i] = "****"
			}
		}
	}
	cleaned := strings.Join(words, " ")
	return cleaned
}