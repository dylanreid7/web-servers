package main

import (
	"github.com/dylanreid7/web-servers/internal/database"
)

func (cfg *apiConfig) handlerGetChirps(w http.ResponseWriter, r *http.Request) {
	chirps, err := database.GetChirps()

}