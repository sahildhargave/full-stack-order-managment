package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()

}
