package response

import "net/http"

func writeWithStausCode(w http.ResponseWriter, statuscode int, b []byte) {
	w.WriteHeader(statuscode)
	w.Write(b)
}
