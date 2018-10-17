package enhance

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"api/session"
)

type EnhanceRouter struct {
	router *httprouter.Router
}

func NewEnhanceRouter(r *httprouter.Router) *EnhanceRouter {

	return &EnhanceRouter{
		router: r,
	}
}

func (e EnhanceRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var isLogined string = "y"

	sessionid, err := r.Cookie("X-SESSIONID")

	if err != nil || session.CheckSession(sessionid.Value) {
	    isLogined = "n"
	}

	http.SetCookie(w, &http.Cookie {
		Name: "X-Logined",
		Value: isLogined,
	})

	e.router.ServeHTTP(w, r)
}


