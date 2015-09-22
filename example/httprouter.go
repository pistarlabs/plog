package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pistarlabs/plog"
)

func hello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello world!")
}

func main() {
	router := httprouter.New()
	router.GET("/", hello)

	plog := plog.Default()

	http.ListenAndServe(":8080", plog.Handler(router))
}
