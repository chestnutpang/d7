package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)


func main(){
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}


func indexHandler(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(os.Stdout, "URL.Path = %q\n", req.URL.Path)
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request){
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

