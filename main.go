package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", hello)

	if err := http.ListenAndServe(":4000", nil); err != nil {
		fmt.Println("server error: ", err)
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/plain")
	fmt.Fprintf(w, "hello world.\nthe time is: %s", time.Now().Format("2006-01-02 15:04:05"))
}
