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
		lang, ok := langs[c.DefaultQuery("lang", "en")]
		if !ok {
			c.HTML(http.StatusNotFound, "404.tmpl", gin.H{"error": "language not found"})
		}
		c.HTML(http.StatusOK, "index", lang.Home)
	})

	r.GET("/about", func(c *gin.Context) {
		lang, ok := langs[c.DefaultQuery("lang", "en")]
		if !ok {
			c.HTML(http.StatusNotFound, "404.tmpl", gin.H{"error": "language not found"})
		}
		c.HTML(http.StatusOK, "about", lang.About)
	})

	r.Run()
}

// default values in case a string is empty
func def(str, def string) string {
	if str != "" {
		return str
	}
	return def
}
