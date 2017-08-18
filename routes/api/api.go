package api

import (
	"github.com/Fisado/Gofus/routes/api/v1"
	"github.com/nmaggioni/goat"
)

func Init(r *goat.Router) {
	v1.Init(r.Subrouter("/v1"))

	r.Get("/", "", r.IndexHandler)
}
