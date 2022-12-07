package routers

import (
	"github.com/gorilla/mux"
	"github.com/ilies-a/artom-api/app/handlers"
)

func initHomeRouter(r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
}
