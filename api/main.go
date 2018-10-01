package main

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "./handler"
    "fmt"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

    router.GET("/user", handler.CreateUser)

	return router
}

func main() {

	r := RegisterHandlers()

    fmt.Println("listening")

	http.ListenAndServe(":8080", r)
}