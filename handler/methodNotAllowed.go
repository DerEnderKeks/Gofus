package handler

import (
	"github.com/abiosoft/river"
	"github.com/Fisado/Gofus/response"
	"net/http"
)

func MethodNotAllowed(c * river.Context) {
	response.WriteError(c, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
}
