package auth

import (
	"net/http"
	"github.com/Fisado/Gofus/response"
	"strings"
)

func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/api") {
			h.ServeHTTP(w, r)
			return
		}
		if false {
			response.WriteError(w, http.StatusUnauthorized, "lol")
		} else {
		}
	})
}