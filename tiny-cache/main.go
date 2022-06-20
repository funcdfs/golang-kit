package main

/*
$ curl http://localhost:8080/_geecache/scores/Tom
630
$ curl http://localhost:8080/_geecache/scores/kkk
kkk not exist
*/

import (
	"fmt"
	"log"
	"net/http"
	"tiny-cache/tinyCache"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	tinyCache.NewGroup("scores", 2<<10, tinyCache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte("get: " + v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	addr := "localhost:8080"
	peers := tinyCache.NewHTTPPool(addr)
	log.Println("tinyCache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}