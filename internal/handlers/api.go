package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type RespData struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func InfoApi(rw http.ResponseWriter, r *http.Request) {
	respData := RespData{
		Message: "Welcome to the API",
		Time:    time.Now(),
	}
	resp, err := json.Marshal(respData)
	if err != nil {
		log.Println(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(resp)
}
