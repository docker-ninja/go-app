package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/fatih/color"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {
	addr := ":8080"
	http.HandleFunc("/", hello)
	color.Green(fmt.Sprintf("%s - listen on %s", time.Now().Format("01-02-2006 15:04:05"), addr))
	color.Red(http.ListenAndServe(addr, nil).Error())
}
