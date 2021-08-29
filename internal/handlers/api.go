package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
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
		sendErrorJsonResponse(rw, []byte("Internal Server Error"), http.StatusInternalServerError)
		return
	}
	sendJsonResponse(rw, resp, http.StatusOK)
}

// WorkdayDate shows whether date is workday
func WorkdayDate(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var date time.Time
	countryCode := vars["cc"]
	dateStr := vars["date"]

	if dateStr == "today" {
		date = time.Now()
		dateStr = date.Format("2006.01.02")
	} else {
		matched, err := regexp.MatchString(`\d{4}-\d{2}-\d{2}`, dateStr)
		if err != nil {
			log.Println(err)
			sendErrorJsonResponse(rw, []byte("Internal Server Error"), http.StatusInternalServerError)
			return
		}
		if !matched {
			sendErrorJsonResponse(rw, []byte("Wrong date format - required YYYY-MM-DD"), http.StatusBadRequest)
			return
		}
		dateStr = strings.Replace(dateStr, "-", ".", 2)
	}

	isWorkday, err := IsDateWorkday(dateStr, countryCode)
	if err != nil {
		log.Println(err)
		sendErrorJsonResponse(rw, []byte("No data found"), http.StatusNotFound)
		return
	}
	respData := StandartResponseData{
		DateString: dateStr,
		IsWorkday:  isWorkday,
	}

	resp, err := json.Marshal(respData)
	if err != nil {
		log.Println(err)
		sendErrorJsonResponse(rw, []byte("Internal Server Error"), http.StatusInternalServerError)
		return
	}
	sendJsonResponse(rw, resp, http.StatusOK)
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
