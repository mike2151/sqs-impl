package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const PORT = 1801

func messageHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Got your request!")
}

func main() {
	http.HandleFunc("/", messageHandler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), nil))
}
