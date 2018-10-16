package handler

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
)

func UserRegister(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	fmt.Fprintln(w, "Hello UserRegister")
}

func UserLogin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

    fmt.Fprintln(w,"This is UserLogin, Hello ", p.ByName("username"))
}