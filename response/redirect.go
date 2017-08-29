package response

import "net/http"

func RedirectTemporary(w http.ResponseWriter, target string) {
	redirect(w, http.StatusFound, target)
}

func RedirectPermanent(w http.ResponseWriter, target string)  {
	redirect(w, http.StatusMovedPermanently, target)
}

func redirect(w http.ResponseWriter, statuscode int, target string) {
	w.Header().Set("Location", target)
	w.WriteHeader(statuscode)
}