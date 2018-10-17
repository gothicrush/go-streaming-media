package router

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func UploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	ck,_  := r.Cookie("X-Logined")
	if ck.Value == "n" {
		fmt.Println("UploadHandler：没有登录")
		http.Redirect(w, r, "/login", 302)
	} else if ck.Value == "y" {
		fmt.Println("UploadHandler：已经登录")
	}
}