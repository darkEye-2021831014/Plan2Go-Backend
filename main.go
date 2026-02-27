package main

import (
	"net/http"
	"plan2go-backend/cmd"
)

func main() {
	cmd.Serve()
	http.HandleFunc("/", homeHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Go backend is running 🚀"))
}
