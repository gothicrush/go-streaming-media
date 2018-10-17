package router

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	ck,_  := r.Cookie("X-Logined")
	if ck.Value == "n" {
		fmt.Println("LoginHandler：没有登录")
	} else if ck.Value == "y" {
		fmt.Println("LoginHandler：已经登录")
		http.Redirect(w, r, "/index", 302)
	}

	templateEngine.ExecuteTemplate(w, "login.html", "v")
}
