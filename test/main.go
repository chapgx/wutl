package main

import (
	"fmt"
	"net/http"

	"github.com/chapgx/wutl"
)

func main() {
	mx := http.NewServeMux()

	mx.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("made it to main")
		w.Write([]byte("Hellow"))
	})

	h := wutl.NewHandler(mx)
	h.AddMiddleware(
		log1,
		log2,
		log3,
	)

	server := http.Server{Addr: ":8080", Handler: h}
	fmt.Println("server in 8080")
	if e := server.ListenAndServe(); e != nil {
		panic(e)
	}
}

func log1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("log 1")
		next.ServeHTTP(w, r)
	})
}

func log2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("log 2")
		next.ServeHTTP(w, r)
	})
}

func log3(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("log 3")
		next.ServeHTTP(w, r)
	})
}
