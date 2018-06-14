package main

import (
	"fmt"
	"net/http"
	"time"
)

func motd() {
	t0 := time.Now()
	fmt.Println("welcome! okaq sakkei web server runnin on localhost:8080")
	fmt.Printf("%s\n", t0.String())
}

func SekkeiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,"sekkei.html")
}

func main() {
	motd()
	http.HandleFunc("/", SekkeiHandler)
	http.ListenAndServe(":8080", nil)
}

