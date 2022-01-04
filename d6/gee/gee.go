package gee


import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

type HandlerFunc func(*Context)


type (
	RouterGroup struct {
		prefix      string
		middlewares []HandlerFunc
		parent      *RouterGroup
		engine      *Engine
	}

	Engine struct {
		*RouterGroup
		router       *router
		groups       []*RouterGroup
		htmlTemplate *template.Template
		funcMap      template.FuncMap
	}
)


func New() *Engine {
	engine := &Engine{router: newRouter()}
}