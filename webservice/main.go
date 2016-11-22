package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
}

func hello(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")

	if len(name) < 2 {
		sendResponse(http.StatusBadRequest, "The name parameter must be at least 2 characters long", w)
	} else {
		body := fmt.Sprintf("hello %s", r.FormValue("name"))

		sendResponse(http.StatusOK, body, w)
	}
}

func sendResponse(statusCode int, responseBody string, w http.ResponseWriter) {
	res := Response{
		Code:   statusCode,
		Result: responseBody,
	}

	json, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func main() {
	var serverPort int
	flag.IntVar(&serverPort, "port", 8000, "The port to run the server on")
	flag.Parse()

	// Add ads the function thats going to handle that response
	http.HandleFunc("/", hello)

	// Starts the web server
	// The first argument in this method is the port you want your server to run on
	// The second is a handler. However we have already added this in the line above so we just pass in nil
	http.ListenAndServe(fmt.Sprintf(":%v", serverPort), nil)
}
