package main

import "net/http"

func main() {
	sm := http.NewServeMux()
	sm.Handle("/health", http.HandlerFunc(healthHandler))

	server := &http.Server{
		Addr:    ":8080",
		Handler: sm,
	}

	server.ListenAndServe()
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HEALTHY"))
}
