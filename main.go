package main

import (
	"fmt"
	"net/http"
	"time"
)

const USERNAME = "test"
const PASSWORD = "secret"

func main() {
    mux := http.DefaultServeMux

    var handler http.Handler = mux
    handler = MiddlewareAuth(handler)
    
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/hello", handlerHello)
	http.HandleFunc("/", handlerIndex)
    http.HandleFunc("/users", handlerUser)
    http.HandleFunc("/json", handlerResponseJson)
	http.HandleFunc("/index", handlerIndex)
    http.HandleFunc("/cool", handlerCool)

	var address = ":9000"
	fmt.Printf("server started at http://localhost%s\n", address)

	server := new(http.Server)
	server.Addr = address
	server.ReadTimeout = time.Second * 10
	server.WriteTimeout = time.Second * 10
    server.Handler = handler

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
