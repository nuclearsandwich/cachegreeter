package main

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"log"
	"net/http"
	"os"
)

var cache *memcache.Client

func greetHandler(w http.ResponseWriter, request *http.Request) {
	item, err := cache.Get("Name")
	if err != nil {
		log.Fatal(err)
	}
	name := string(item.Value)
	fmt.Fprintf(w, "Hello, %s!", name)
}

func nameHandler(w http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	param, hasparam := query["name"]
	var name string
	if len(param) != 1 {
		name = "World"
	} else {
		name = param[0]
	}
	if hasparam {
		cache.Set(&memcache.Item{Key: "Name", Value: []byte(name), Flags: 0, Expiration: 0})
	}
	fmt.Fprintf(w, "Stored name '%s'", name)

}


func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8090"
	}
	memcache_url := os.Getenv("MEMCACHE_URL")
	if memcache_url == "" {
		memcache_url = "127.0.0.1:11211"
	}
	cache = memcache.New(memcache_url)
	cache.Set(&memcache.Item{Key: "Name", Value: []byte("World"), Flags: 0, Expiration: 0})
	http.HandleFunc("/", greetHandler)
	http.HandleFunc("/name", nameHandler)
	log.Fatal(http.ListenAndServe(port, nil))
}
