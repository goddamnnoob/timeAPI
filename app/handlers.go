package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func GetTime(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	queryValues := r.URL.Query()
	if !queryValues.Has("tz") {
		fmt.Fprintln(rw, time.Now())
	} else {
		tz := queryValues.Get("tz")
		if tz == "" {
			fmt.Fprintln(rw, time.Now())
		} else {
			loc, err := time.LoadLocation(tz)
			if err != nil {
				fmt.Fprintln(rw, "Invalid Input 404 error")
			} else {
				fmt.Fprintln(rw, tz+": "+time.Now().In(loc).String())
			}
		}
	}
}
