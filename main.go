package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")

	langs := languages()

	r.LoadHTMLGlob("templates/**")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home", useLang(langs, c))
	})
	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about", useLang(langs, c))
	})

	r.Run()
}

// helper function to make the code more DRY. It selects the language from the
// context or fails the request.
func useLang(langs map[string]language, c *gin.Context) language {
	reqLang := c.DefaultQuery("lang", "en")
	lang, ok := langs[reqLang]
	if !ok {
		c.HTML(http.StatusNotFound, "404.tmpl", gin.H{"error": "language not found"})
	}
	lang.Header.Lang = reqLang
	return lang
}
