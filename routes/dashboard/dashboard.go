package dashboard

import (
	"github.com/abiosoft/river"
	"os"
	"io/ioutil"
	"github.com/Fisado/Gofus/response"
	"mime"
	"path/filepath"
	"net/http"
)

func Init(e *river.Endpoint) {
	e.Get("/*filepath", handleStatic)
}

func handleStatic(c *river.Context) {
	wd, _ := os.Getwd()

	filePath := c.Param("filepath")
	if filePath == "/" {
		filePath += "index.html"
	}

	localFilePath := filepath.Join(wd, "/static", filePath)

	b, err := ioutil.ReadFile(localFilePath)
	if err != nil {
		response.WriteError(c, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	response.WriteFile(c, b, mime.TypeByExtension(filepath.Ext(localFilePath)), "")
}
