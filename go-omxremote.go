package main

import (
	"log"
	"net/http"
)

var p Player

func main() {
	parseCmdParam()

	router := initRouter()

	log.Printf("http://localhost%s on %s", bindAddr, videosPath)
	log.Fatal(http.ListenAndServe(bindAddr, router))
}
