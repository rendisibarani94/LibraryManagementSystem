package helpers

import (
	"encoding/json"
	"net/http"
)

type ResponseWithData struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type ResponseWithoutData struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

func Response(w http.ResponseWriter, code int, message string, data interface{}){
	w.Header().Set("Content-Type", "application/json") // Setting header "Content-Type", "application/json"
	w.WriteHeader(code) // 

	var response interface{} // interface{} or any
	status := "success"
	
	// more code condition, better
	if code >= 400 { // just checking one single code condition
		status = "failed"
	}

	if data != nil {
		response = &ResponseWithData{
			Status: status,
			Message: message,
			Data: data,
		}
	}else {
		response = &ResponseWithoutData{
			Status: status,
			Message: message,
		}
	}

	res, _ := json.Marshal(response) // marshall / construct into json form 
	w.Write(res) // writing response
}