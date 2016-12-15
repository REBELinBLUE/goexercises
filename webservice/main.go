package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/REBELinBLUE/goexercises/webservice/developers"
	"github.com/REBELinBLUE/goexercises/webservice/response"
	"github.com/julienschmidt/httprouter"
)

func main() {
	var serverPort int
	flag.IntVar(&serverPort, "port", 8000, "The port to run the server on")
	flag.Parse()

	router := httprouter.New()

	router.GET("/", hello)
	router.POST("/developers", developers.AddDeveloper)
	router.GET("/developers", developers.ShowDevelopers)
	router.GET("/developers/:developer", developers.ShowDeveloper)

	http.ListenAndServe(fmt.Sprintf(":%v", serverPort), router)
}

func hello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	name := r.FormValue("name")

	if len(name) < 2 {
		response.Send(http.StatusBadRequest, "The name parameter must be at least 2 characters long", w)
	} else {
		body := fmt.Sprintf("hello %s", r.FormValue("name"))

		response.Send(http.StatusOK, body, w)
	}
}
