package main


import (
	"net/http"
	"gee"
)


func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello world!")
	})

	r.GET("/panic", func(c *gee.Context) {
		name := []string{"link"}
		c.String(http.StatusOK, name[100])
	})

	r.Run(":9999")
}