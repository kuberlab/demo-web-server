package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8082",nil)
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	v := os.Getenv("SERVER_VERSION")
	w.Write([]byte("Hello World!!! I'm "+v+"."))
}