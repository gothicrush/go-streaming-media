package handler

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "io"
    "fmt"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    fmt.Println("???")

	io.WriteString(w, "CreateUser Handler")
}