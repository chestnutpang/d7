package main


import (
	"fmt"
	"geecache"
	"log"
	"net/http"
)


var db = map[string]string {
	"Tom": "630",
	"Jack": "576",
	"Sam": "657",
}

func main() {
	geecache.NewGroup("scores", 2 << 10, geecache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		},
	))
	addr := "localhost:9999"
	peers := geecache.NewHTTPPool(addr)
	log.Println("geecache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}