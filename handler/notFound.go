package handler

import (
	"github.com/abiosoft/river"
	"github.com/Fisado/Gofus/response"
	"net/http"
	"strings"
	"fmt"
)

func NotFound(c *river.Context) {
	shortName := strings.Replace(c.URL.Path, "/", "", 1)
	if len(shortName) < 8 || len(shortName) > 256 {
		notFound(c)
		return
	}

	shortName = strings.Split(shortName, ".")[0]
	fmt.Println(shortName)

	// TODO
}

func notFound(c *river.Context) {
	response.WriteError(c, http.StatusNotFound, http.StatusText(http.StatusNotFound))
}
