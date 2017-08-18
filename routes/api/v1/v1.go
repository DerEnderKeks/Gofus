package v1

import (
	"github.com/nmaggioni/goat"
)

func Init(r *goat.Router) {
	r.Get("/", "", r.IndexHandler)
	r.Get("/teapot", "", Teapot)
}
