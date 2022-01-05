package main


import (
	"fmt"
	"html/template"
	"time"
	"gee"
)

type student struct {
	Name string
	Age int8
}

func FormatAsData(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprint("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsData": FormatAsData,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{}
}