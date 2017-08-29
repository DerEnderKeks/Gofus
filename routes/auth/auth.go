package auth

import (
	"net/http"
	"github.com/Fisado/Gofus/response"
	"github.com/nmaggioni/goat"
	"github.com/abiosoft/river"
)

func Init(e *river.Endpoint) {
	e.Get("/session", SessionAuth)
}

func SessionAuth(w http.ResponseWriter, r *http.Request, p goat.Params) {
	response.WriteJSON(w, map[string]string {
		"TO": "DO", // TODO
	})
}
