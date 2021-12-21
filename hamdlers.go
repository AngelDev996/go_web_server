package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hellowda")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the API endpint")
}
