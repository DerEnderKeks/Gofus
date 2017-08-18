package response

import (
	"net/http"
	"encoding/json"
)


func WriteError(w http.ResponseWriter, statuscode int, err string) {
	WriteJSONWithStatus(w, statuscode, map[string]string{
		"error": err,
	})
}

func WriteJSON(w http.ResponseWriter, v interface{}) error {
	return writeJSONWithStatus(w, http.StatusOK, v)
}

func WriteJSONWithStatus(w http.ResponseWriter, statuscode int, v interface{}) error {
	return writeJSONWithStatus(w, statuscode, v)
}

func writeJSONWithStatus(w http.ResponseWriter, statuscode int, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	writeWithStausCode(w, statuscode, b)
	return nil
}

