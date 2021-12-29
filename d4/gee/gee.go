package gee

import (
	"log"
	"net/http"
)

type HandlerFunc func(*Context)

type (
	RouterGroup struct {
		prefix     string
		middleware []HandlerFunc
		parent     *RouterGroup
		engine     *Engine
	}

	Engine struct {
		*RouterGroup
		router *router
		groups []*RouterGroup
	}
)

func New() *Engine {
	engine := &Engine{router: newRouter()}
	eninge.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}


// 创建一个新的路由分组
// 所有分组分享一个 Engine 实例
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup {
		prefix: group.prefix + prefix
		parent: group
		engine: engine
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}


func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.addRoute(method, pattern, handler)
}


fun (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

fun (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
