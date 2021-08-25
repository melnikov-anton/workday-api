package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type RespData struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

type StandartResponseData struct {
	DateString string `json:"date"`
	IsWorkday  bool   `json:"is_workday"`
}

// InfoApi shows info about API
func InfoApi(rw http.ResponseWriter, r *http.Request) {
	respData := RespData{
		Message: "Welcome to the Workday API",
		Time:    time.Now().Format("2006-01-02 15:04:05"),
	}
	resp, err := json.Marshal(respData)
	if err != nil {
		log.Println(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(resp)
}

// WorkdayToday shows whether today is workday
func WorkdayToday(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	countryCode := vars["cc"]
	today := time.Now()
	todayString := today.Format("2006.01.02")
	isWorkday, err := IsDateWorkday(todayString, countryCode)
	if err != nil {
		log.Println(err)
	}
	respData := StandartResponseData{
		DateString: todayString,
		IsWorkday:  isWorkday,
	}

	resp, err := json.Marshal(respData)
	if err != nil {
		log.Println(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(resp)
}

func IsDateWorkday(dateStr, countryCode string) (bool, error) {
	currentYearStr := strconv.Itoa(time.Now().Year())
	fileName := filepath.Join(appConfig.AppDataDir, countryCode, currentYearStr+".txt")
	if fileExists(fileName) {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			return false, err
		}
		return !strings.Contains(string(data), dateStr), nil
	} else {
		return false, errors.New("no data for a year or country")
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
