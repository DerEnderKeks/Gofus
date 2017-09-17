package v1

import (
	"github.com/Fisado/Gofus/response"
	"net/http"
	"github.com/abiosoft/river"
	"regexp"
)

func Teapot(c *river.Context) {
	r, _ := regexp.Compile("\n\\s*")
	html := r.ReplaceAllString(`<html>
		<head>
		<style>
		body {
			all: unset;
			overflow: hidden;
			background-color: black;
		}
		@keyframes spin {
			from {transform:rotate(0deg);}
			to {transform:rotate(360deg);}
		}
		</style>
		</head>
		<body>
		<div>
		<iframe style="animation: spin infinite 5s linear;" width="100%" height="100%" src="https://www.youtube.com/embed/8E6_PNxed5Y?autoplay=1&loop=1&modestbranding=1&showinfo=0&rel=0&disablekb=1&controls=0" frameborder="0" allowfullscreen></iframe>
		</div>
		</body>
		</html>`, "")
	response.WriteHTMLWithStatus(c, http.StatusTeapot, html)
}