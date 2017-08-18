package auth

import (
	"net/http"
	"github.com/Fisado/Gofus/response"
	"github.com/nmaggioni/goat"
)

func Auth(w http.ResponseWriter, r *http.Request, p goat.Params) {
	response.WriteJSON(w, map[string]string {
		"TO": "DO", // TODO
	})
}
