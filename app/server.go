package app

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Start() {
	router := httprouter.New()
	router.GET("/api/time", GetTime)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
