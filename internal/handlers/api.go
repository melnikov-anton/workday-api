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

const (
	STATUS_INTERNAL_SERVER_ERROR = "Internal Server Error"
	MESSAGE_WRONG_DATE_FORMAT    = "wrong date format"
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
		sendErrorJsonResponse(rw, []byte(STATUS_INTERNAL_SERVER_ERROR), http.StatusInternalServerError)
		return
	}
	sendJsonResponse(rw, resp, http.StatusOK)
}

// WorkdayDate shows whether date is workday
func WorkdayDate(rw http.ResponseWriter, r *http.Request) {

	dateStr, countryCode, err := getDateAndCountryCode(r)
	if err != nil {
		log.Println(err)
		switch err.Error() {
		case MESSAGE_WRONG_DATE_FORMAT:
			sendErrorJsonResponse(rw, []byte("Wrong date format - required YYYY-MM-DD"), http.StatusBadRequest)
		default:
			sendErrorJsonResponse(rw, []byte(STATUS_INTERNAL_SERVER_ERROR), http.StatusInternalServerError)
		}
		return
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
		sendErrorJsonResponse(rw, []byte(STATUS_INTERNAL_SERVER_ERROR), http.StatusInternalServerError)
		return
	}
	sendJsonResponse(rw, resp, http.StatusOK)
}

// WorkdayDateSimple returns simple response (1 = workday, 0 = holyday or error code)
func WorkdayDateSimple(rw http.ResponseWriter, r *http.Request) {
	dateStr, countryCode, err := getDateAndCountryCode(r)
	if err != nil {
		log.Println(err)
		switch err.Error() {
		case MESSAGE_WRONG_DATE_FORMAT:
			sendSimpleResponse(rw, []byte("400"), http.StatusBadRequest)
		default:
			sendSimpleResponse(rw, []byte("500"), http.StatusInternalServerError)
		}
		return
	}

	isWorkday, err := IsDateWorkday(dateStr, countryCode)
	if err != nil {
		log.Println(err)
		sendSimpleResponse(rw, []byte("404"), http.StatusNotFound)
		return
	}
	if isWorkday {
		sendSimpleResponse(rw, []byte("1"), http.StatusOK)
		return
	} else {
		sendSimpleResponse(rw, []byte("0"), http.StatusOK)
		return
	}
}

func getDateAndCountryCode(r *http.Request) (string, string, error) {
	vars := mux.Vars(r)
	var date time.Time
	countryCode := vars["cc"]
	dateStr := vars["date"]

	if dateStr == "today" {
		date = time.Now()
		dateStr = date.Format("2006.01.02")
		return dateStr, countryCode, nil
	} else {
		matched, err := regexp.MatchString(`\d{4}-\d{2}-\d{2}`, dateStr)
		if err != nil {
			log.Println(err)
			return "", "", errors.New("internal server error")
		}
		if !matched {
			return "", "", errors.New(MESSAGE_WRONG_DATE_FORMAT)
		}
		dateStr = strings.Replace(dateStr, "-", ".", 2)
		return dateStr, countryCode, nil
	}
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
