package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Helo service that says hi to the world
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "world"
		}
		message := fmt.Sprintf("Hello, %s!", name)

		fmt.Fprint(w, message)
	})

	fmt.Println("Starting on :8080")
	http.ListenAndServe(":8080", nil)
}
