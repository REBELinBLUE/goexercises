package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// ResponseContent represents the expected structure of the response
type ResponseContent struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
}

// Send the HTTP response
func Send(statusCode int, responseBody interface{}, w http.ResponseWriter) {
	res := ResponseContent{
		Code:   statusCode,
		Result: responseBody,
	}

	json, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(json)
}
