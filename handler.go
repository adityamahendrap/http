package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Welcome"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello World!"
	w.Write([]byte(message))
}

func handlerCool(w http.ResponseWriter, r *http.Request) {
    var filepath = path.Join("views", "index.html")
    templ, err := template.ParseFiles(filepath)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var data = map[string]interface{} {
        "title": "Learning Golang Web",
        "name":  "Batman",
    }

    err = templ.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func handlerUser(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case "GET":
            w.Write([]byte("get"))
        case "POST":
            w.Write([]byte("post"))
        default:
            http.Error(w, "", http.StatusBadRequest)
    }
}

func handlerResponseJson(w http.ResponseWriter, r *http.Request) {
    data := []struct {
        Name string
        Age  int
    } {
        { "Richard Grayson", 24 },
        { "Jason Todd", 23 },
        { "Tim Drake", 22 },
        { "Damian Wayne", 21 },
    }

    jsonByte, err := json.Marshal(data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Conttent-Type", "application/json")
    w.Write(jsonByte)
}