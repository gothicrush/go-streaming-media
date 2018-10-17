package router

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func SignoutHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	ck,_  := r.Cookie("X-Logined")
	if ck.Value == "n" {
		fmt.Println("SignoutHandler：没有登录")
		http.Redirect(w, r, "/index", 302)
	} else if ck.Value == "y" {
		fmt.Println("SignoutHandler：已经登录")
	}
}
