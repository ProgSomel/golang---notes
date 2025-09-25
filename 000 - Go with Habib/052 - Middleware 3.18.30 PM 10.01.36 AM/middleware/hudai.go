package middleware

import (
	"log"
	"net/http"
)

func Hudai(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		log.Println("Hi, I am Hudai")
		next.ServeHTTP(w, r)
	})
}