package main

import (
	"fmt"
	"log"
	"net/http"
)

type APIServer struct {
	address string
}

func NewAPIServer(address string) *APIServer {
	return &APIServer{
		address,
	}
}

func (s *APIServer) Run() {

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World: " + request.RequestURI))
	})

	router.Handle("GET /users/{userId}", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		userId := request.PathValue("userId")

		writer.Write([]byte("userId: " + userId))
	}))

	middlewareChain := combineMiddleware(
		Auth,
		Logger,
	)

	server := http.Server{
		Addr:    s.address,
		Handler: middlewareChain(router),
	}

	fmt.Println("API server listening on " + s.address)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func Logger(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	}
}

func Auth(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Auth")
		next.ServeHTTP(w, r)
	}
}

type Middleware func(http.Handler) http.HandlerFunc

func combineMiddleware(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {

		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}

		return next.ServeHTTP
	}
}
