package v1

import (
	"github.com/abiosoft/river"
)

func Init(e *river.Endpoint) {
	e.Get("/teapot", Teapot)
}
