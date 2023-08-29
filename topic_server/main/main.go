package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/topic/", topicserver.HandlerTopic)
	log.Fatal(http.ListenAndServe(":2017", nil))
}
