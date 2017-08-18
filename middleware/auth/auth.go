package auth

import (
	"net/http"
	"github.com/Fisado/Gofus/response"
)

func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if false {
			response.WriteError(w, http.StatusUnauthorized, "lol")
		} else {
			h.ServeHTTP(w, r)
		}
	})
}