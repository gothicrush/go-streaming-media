package api

import (
	"api/handler"
	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/user", handler.UserRegister)

	router.POST("/user/:username", handler.UserLogin)

	return router
}
