package main

import (
	"io"
	"net/http"
)

const MaxClientCount = 100

func main(){

	helloWorldHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		io.WriteString(w, "hello world!")
	})

	http.Handle("/", maxClients(helloWorldHandler, MaxClientCount))
	http.ListenAndServe(":5011", nil)
}

func maxClients(h http.Handler, n int) http.Handler{
	sema := make(chan struct{}, n)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		sema <- struct{}{}
		defer func(){ <- sema }()
		h.ServeHTTP(w, r)
	})
}

