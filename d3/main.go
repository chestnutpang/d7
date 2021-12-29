package main
import (
	"net/http"
	"gee"
	"log"
)


func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello World.</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "\nhello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *gee.Context) {
		c.String(http.StatusOK, "\nhello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	r.GET("/assets/*filename", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filename": c.Param("filename")})
	})

	log.Println("server start at 0.0.0.0:9999")
	log.Fatal(r.Run(":9999"))
}