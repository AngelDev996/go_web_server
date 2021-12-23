package main

import (
	"fmt"
	"net/http"
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true //Si es falso no deja pasar el handler
			fmt.Println("Checking Auth")
			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}
