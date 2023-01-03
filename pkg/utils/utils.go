package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type StatusStruct struct {
	temp     string
	mode     string
	fanSpeed string
	on       string
}

type State struct {
	Mode   string
	Stemp  string
	Shum   string
	F_rate string
	F_dir  string
	Power  bool
}

func SendRequest(airconIp string, power string, mode string, temp string, fanRate string) {
	_, err := http.Get(fmt.Sprintf("http://10.0.0.24/aircon/set_control_info?pow=%s&mode=%s&stemp=%s&shum=0&f_rate=%s&f_dir=3", power, mode, temp, fanRate))

	if err != nil {
		log.Fatal("Error toggling aircon", err)
	}
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

func StoreCurrentState(airconIp string) State {
	var power bool
	var mode string
	var stemp string
	var shum string
	var f_rate string
	var f_dir string

	powerMap := map[string]bool{
		"1": true,
		"0": false,
	}

	for _, value := range getCurrentStatus(airconIp) {
		if strings.Contains(value, "pow") {
			power = powerMap[strings.Split(value, "=")[1]]
		}

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

	finalState := State{mode, stemp, shum, f_rate, f_dir, power}

	return finalState
}

func CurrentStatus(airconIp string) {
	var temp string
	var mode string
	var fanSpeed string
	var power string

	modeMap := map[string]string{
		"3": "Cold",
		"4": "Heat",
	}

	powerMap := map[string]string{
		"1": "On",
		"0": "Off",
	}

	fanMap := map[string]string{
		"B": "Night",
		"3": "1",
		"4": "2",
		"5": "3",
		"6": "4",
		"7": "5",
	}

	for _, value := range getCurrentStatus(airconIp) {
		if strings.Contains(value, "stemp") {
			temp = strings.Split(value, "=")[1]
		}

		if strings.Contains(value, "mode") {
			mode = modeMap[strings.Split(value, "=")[1]]
		}

		if strings.Contains(value, "pow") {
			power = powerMap[strings.Split(value, "=")[1]]
		}

		if strings.Contains(value, "f_rate") {
			fanSpeed = fanMap[strings.Split(value, "=")[1]]
		}
	}

	fmt.Println(StatusStruct{temp, mode, fanSpeed, power})
}
