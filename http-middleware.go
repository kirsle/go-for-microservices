package main

import (
	"fmt"
	"net/http"
)

func myLoggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("GET %s\n", r.URL)
		next.ServeHTTP(w, r)
	})
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func main() {
	myHandler := http.HandlerFunc(myHandler)

	// Middlewares are applied by wrapping functions
	http.Handle("/", myLoggingHandler(myHandler))
	http.ListenAndServe(":8000", nil)
}
