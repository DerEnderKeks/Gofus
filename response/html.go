package response

import (
	"net/http"
)

func WriteHTML(w http.ResponseWriter, html string) error {
	return writeHTMLWithStatus(w, http.StatusOK, html)
}

func WriteHTMLWithStatus(w http.ResponseWriter, statuscode int, html string) error {
	return writeHTMLWithStatus(w, statuscode, html)
}

func writeHTMLWithStatus(w http.ResponseWriter, statuscode int, html string) error {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	writeWithStausCode(w, statuscode, []byte(html))
	return nil
}