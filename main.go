package main

import (
	"api"
	"fmt"
	"net/http"
)

func main() {

	r := api.RegisterHandlers()

    fmt.Println("listening")

	http.ListenAndServe(":8080", r)
}