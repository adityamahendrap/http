package main

import (
	"fmt"
	"net/http"
	"time"
)

type CustomMux struct {
	http.ServeMux
	middlewares []func(next http.Handler) http.Handler
}

func (c *CustomMux) RegisterMiddleware(next func(next http.Handler) http.Handler) {
	c.middlewares = append(c.middlewares, next)
}

func (c *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServeMux

	for _, next := range c.middlewares {
		current = next(current)
	}

	current.ServeHTTP(w, r)
}

const USERNAME = "test"
const PASSWORD = "secret"

type M map[string]interface{}

var cookieName = "thisIsCookie"

func main() {
	mux := new(CustomMux)

	mux.RegisterMiddleware(MiddlewareAuth)

	mux.HandleFunc("/cookie", handlerCreateCookie)
	mux.HandleFunc("/cookie/destroy", handlerDeleteCookie)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/hello", handlerHello)
	mux.HandleFunc("/", handlerIndex)
	mux.HandleFunc("/users", handlerUser)
	mux.HandleFunc("/json", handlerResponseJson)
	mux.HandleFunc("/index", handlerIndex)
	mux.HandleFunc("/cool", handlerCool)

	var address = ":9000"
	fmt.Printf("server started at http://localhost%s\n", address)

	server := new(http.Server)
	server.Addr = address
	server.ReadTimeout = time.Second * 10
	server.WriteTimeout = time.Second * 10
	server.Handler = mux

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
