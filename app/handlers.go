package app

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

type timeLoc struct {
	Loc  string
	Time string
}

func GetTime(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	queryValues := r.URL.Query()
	rw.Header().Add("Content-Type", "application/json")
	res := []timeLoc{}
	if !queryValues.Has("tz") {
		res = append(res, timeLoc{"LOCAL TIME", time.Now().String()})
	} else {
		tz := queryValues.Get("tz")
		if tz == "" {
			res = append(res, timeLoc{"LOCAL TIME", time.Now().String()})
		} else {
			tzs := strings.Split(tz, ",")
			for i := range tzs {
				loc, err := time.LoadLocation(tzs[i])
				if err != nil {
					res = append(res, timeLoc{"Invalid Input", "404 Error"})
				} else {
					res = append(res, timeLoc{tzs[i], time.Now().In(loc).String()})
				}
			}
		}
	}
	json.NewEncoder(rw).Encode(res)
}
