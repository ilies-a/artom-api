package handlers

import (
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	type Structure struct {
		Message string `json:"message"`
	}

}
