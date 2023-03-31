package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Print("hello nayab")
	mux := http.NewServeMux()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello world</h1>"))
}
