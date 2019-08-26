package main

import (
	"fmt"
	"net/http"
	"time"
)

func HelloWeb(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello web")
}

func TimeConsumeMiddleware(handerFunc http.Handler) http.Handler {
	 return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		handerFunc.ServeHTTP(w, r)

		fmt.Printf("consume %d", time.Since(startTime))
	})
}

func main() {
	http.Handle("/hello/web", TimeConsumeMiddleware(http.HandlerFunc(HelloWeb)))
	http.ListenAndServe(":9999", nil)
}
