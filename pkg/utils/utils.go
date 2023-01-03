package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type State struct {
	Mode   string
	Stemp  string
	Shum   string
	F_rate string
	F_dir  string
}

func getCurrentStatus(airconIp string) []string {
	url := fmt.Sprintf("http://%s/aircon/get_control_info", airconIp)
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("Error setting heat", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)

	return strings.Split(sb, ",")
}

func GetOnState(airconIp string) bool {
	var onState bool
	mapStates := map[string]bool{
		"1": true,
		"0": false,
	}

	for _, value := range getCurrentStatus(airconIp) {
		if strings.Contains(value, "pow") {
			onState = mapStates[strings.Split(value, "=")[1]]
		}
	}

	return onState
}

func GetCurrentTemp(airconIp string) string {
	var currentTemp string

	for _, value := range getCurrentStatus(airconIp) {
		if strings.Contains(value, "stemp") {
			currentTemp = strings.Split(value, "=")[1]
		}
	}

	return currentTemp
}

func GetCurrentMode(airconIp string) string {
	var currentMode string

	for _, value := range getCurrentStatus(airconIp) {
		if strings.Contains(value, "mode") {
			currentMode = strings.Split(value, "=")[1]
		}
	}

	return currentMode
}

func StoreCurrentState(airconIp string) State {
	var mode string
	var stemp string
	var shum string
	var f_rate string
	var f_dir string

	for _, value := range getCurrentStatus(airconIp) {
		if strings.Contains(value, "mode") {
			mode = strings.Split(value, "=")[1]
		}

		if strings.Contains(value, "stemp") {
			stemp = strings.Split(value, "=")[1]
		}

		if strings.Contains(value, "shum") {
			shum = strings.Split(value, "=")[1]
		}

		if strings.Contains(value, "f_rate") {
			f_rate = strings.Split(value, "=")[1]
		}

		if strings.Contains(value, "f_dir") {
			f_dir = strings.Split(value, "=")[1]
		}
	}

	finalState := State{mode, stemp, shum, f_rate, f_dir}

	return finalState
}
