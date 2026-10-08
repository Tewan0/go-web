package main

import (
	"fmt"
	"net/http"
)

func helloGoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Go ^_^ ....")
}

func hiGoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi Go ^_^ ....")
}

func main() {
	http.HandleFunc("/hellogo", helloGoHandler)
	http.HandleFunc("/higo", hiGoHandler)

	fmt.Printf("Server start at port 3030")

	// start server
	if err := http.ListenAndServe(":3030", nil); err != nil {
		fmt.Printf("Server start error: %v", err)
	}
}
