package main

import (
	"fmt"
	"log"
	"net/http"
)


type Engine struct{}


func main(){
	engine := new(Engine)
	log.Println("server start")
	// 实现 ServeHTTP 接口的实例，后续http请求都会通过该实例的 ServeHTTP 函数处理
	log.Fatal(http.ListenAndServe(":9999", engine))
}


// Engine 结构体实现 ServeHTTP 接口
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request){
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
