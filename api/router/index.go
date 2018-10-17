package router

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	ck,_ := r.Cookie("X-Logined")
	if ck.Value == "n" {
		fmt.Println("indexHandler：没有登录")
	} else if ck.Value == "y" {
		fmt.Println("indexHandler：已经登录")
	}

	fmt.Println(ck)

	templateEngine.ExecuteTemplate(w, "index.html", "v")
}
