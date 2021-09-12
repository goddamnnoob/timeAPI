package app

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetTime(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	queryValues := r.URL.Query()
	fmt.Fprintf(rw, "hello"+queryValues.Get("tz"))
}
