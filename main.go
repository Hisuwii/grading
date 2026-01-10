package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Index().Render(r.Context(), w)
	})
	http.ListenAndServe(":8080", nil)
}
