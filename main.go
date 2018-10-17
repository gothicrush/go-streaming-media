package main

import (
	"api/router"
	"api/enhance"
	"net/http"
)

func main() {

	r := router.RegisterHandlers()

	enhanceRouter := enhance.NewEnhanceRouter(r)

	http.ListenAndServe(":8080", enhanceRouter)
}