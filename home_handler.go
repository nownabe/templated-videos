package main

import (
	"fmt"
	"net/http"
)

type homeHandler struct{}

func newHomeHandler() *homeHandler {
	return &homeHandler{}
}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
