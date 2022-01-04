package gee


import (
	"log"
	"net/http"
	"strings"
)

// 请求处理函数结构定义
type HandlerFunc func(*Context)


type (
	RouterGroup struct {
		prefix string
		middlewares []HandlerFunc
		parent *RouterGroup
		engine *Engine
	}

	Engine struct {
		*RouterGroup  // 内嵌结构体 engine作为最顶层分组 拥有RouterGroup所有能力
		router *router
		groups []*RouterGroup  // 存储所有分组
	}
)


func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}


func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// 定义路由分组加入中间件的函数
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}


func (group *RouterGroup) addRoute(method, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}


func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}


func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}


func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	c := newContext(w, req)
	c.handlers = middlewares
	engine.router.handle(c)
}
