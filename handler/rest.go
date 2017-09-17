package handler

import (
	"github.com/abiosoft/river"
	"encoding/json"
	"github.com/Fisado/Gofus/database"
)

// GET /:resource?sort=['<field>','<sorting>']&range=[<start>,<end>]&filter={<field>:<value>}
func GetList(c *river.Context) {
	var (
		resource string

		qSort   string
		qOrder  string
		qRange  [2]int = [2]int{0, 25}
		qFilter interface{}
	)

	resource = c.Param("resource")

	if len(c.Query("sort")) > 0 {
		json.Unmarshal([]byte(c.Query("sort")), qSort)
	}
	if len(c.Query("order")) > 0 {
		json.Unmarshal([]byte(c.Query("order")), qOrder)
	}
	if len(c.Query("range")) > 0 {
		json.Unmarshal([]byte(c.Query("range")), qRange)
	}
	if len(c.Query("filter")) > 0 {
		json.Unmarshal([]byte(c.Query("filter")), qFilter)
	}

	model := database.Get(resource)

	if model == nil {
		notFound(c)
	}

	database.DB.Offset(qRange[0]).Limit(qRange[1]).Order(qSort + " " + qOrder).Find(&model) // TODO: Remove
}

// GET /:resource/:id
func GetOne(c *river.Context) {

}

// POST /:resource
func Create(c *river.Context) {

}

// PUT /:resource/:id
func Update(c *river.Context) {

}

// DELETE /:resource/:id
func Delete(c *river.Context) {

}

// GET /:resource?filter={ids:[<id>,...]}
func GetMany(c *river.Context) {

}

// GET /:resource?filter={<ref_field>:<id>}
func GetManyReference(c *river.Context) {

}
