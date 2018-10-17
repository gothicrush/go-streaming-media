package router

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func WatchHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	fmt.Fprintln(w,"This is Watch")
}
