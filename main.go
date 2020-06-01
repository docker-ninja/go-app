package main

import (
	"fmt"
	"net/http"
	"log"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {
	addr := ":8080"
	http.HandleFunc("/", hello)
    log.Println("listen on", addr)
    log.Fatal( http.ListenAndServe(addr, nil) )
}
