package main

import (
	"log"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func oneHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("one handler")
		w.WriteHeader(http.StatusOK)
	})
}

func twoHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("one handler")
		w.WriteHeader(http.StatusOK)
	})
}

func middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware 1")
		next.ServeHTTP(w, r)
	})
}

func middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware 2")

		// don't call next, as if some error occurred
		w.WriteHeader(http.StatusBadRequest)
	})
}

func handler(handler http.Handler, middleware []Middleware) http.Handler {
	if len(middleware) < 1 {
		return handler
	}

	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}

	return handler
}

func main() {
	port := "8000"

	mux := http.NewServeMux()

	oneMiddleware := []Middleware{
		middleware1,
	}

	twoMiddleware := []Middleware{
		middleware1,
		middleware2,
	}

	mux.Handle("/one", handler(
		oneHandler(),
		oneMiddleware,
	))

	mux.Handle("/two", handler(
		twoHandler(),
		twoMiddleware,
	))

	log.Println("server is listening on port: " + port)
	log.Fatal(http.ListenAndServe(":" + port, mux))
}