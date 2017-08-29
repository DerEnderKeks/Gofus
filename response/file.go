package response

import (
	"net/http"
	"mime"
	"strings"
)

var (
	displayTypes = []string{"text/", "image/", "audio/", "video/", "application/pdf", "application/xml", "application/json", "application/xhtml+xml"}
)

func checkDownload(mimeType string) bool {
	download := true

	for _, v := range displayTypes {
		if strings.HasPrefix(mimeType, v) {
			download = false
		}
	}

	return download
}

func WriteFile(w http.ResponseWriter, b []byte, mimeType, filename string) error {
	mimeType, _, err := mime.ParseMediaType(mimeType)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", mimeType)

	if checkDownload(mimeType) {
		w.Header().Set("Content-Disposition", "attachment; filename=\""+filename+"\"")
	}

	writeWithStausCode(w, http.StatusOK, b)

	return nil
}
