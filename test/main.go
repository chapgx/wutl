package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/chapgx/wutl"
)

func main() {
	mx := http.NewServeMux()

	mx.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		p, _ := filepath.Abs("./www/index.html")
		fmt.Println(p)
		b, e := os.ReadFile(p)
		if e != nil {
			http.NotFound(w, r)
			return
		}
		w.Write(b)
	})

	h := wutl.NewHandler(mx)
	h.AddMiddleware(
		wutl.ServeFiles("./www"),
		log1,
		log2,
		log3,
	)

	h.AddMiddleware(log4)

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

func log4(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("log 4")
		next.ServeHTTP(w, r)
	})
}
