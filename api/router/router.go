package router

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
)

var templateEngine *template.Template

func init() {
	t, err := template.ParseFiles(`./templates/index.html`,
		`./templates/login.html`,`./templates/register.html`)

	if err != nil {
		panic(err)
	}

	templateEngine = t
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/index", IndexHandler)
	router.GET("/login", LoginHandler)
	router.GET("/register", RegisterHandler)
	router.GET("/signout", SignoutHandler)
	router.GET("/upload", UploadHandler)
	router.GET("/watch", WatchHandler)

	return router
}