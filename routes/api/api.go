package api

import (
	"github.com/Fisado/Gofus/routes/api/v1"
	"github.com/abiosoft/river"
)

func Init(e *river.Endpoint, r *river.River) {
	v1Endpoint := river.NewEndpoint()
	v1.Init(v1Endpoint)

	r.Handle("/api/v1", v1Endpoint)
}
