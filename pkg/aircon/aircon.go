package aircon

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetOnState() bool {
	mapStates := map[string]bool{
		"1": true,
		"0": false,
	}

	var onState string
	resp, err := http.Get("http://10.0.0.24/aircon/get_control_info")

	if err != nil {
		log.Fatal("Error setting heat", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)

	items := strings.Split(sb, ",")

	for _, item := range items {
		if strings.Contains(item, "pow") {
			onState = strings.Split(item, "=")[1]
		}
	}

	return mapStates[onState]
}

func getCurrentState() string {
	resp, err := http.Get("http://10.0.0.24/aircon/get_control_info")

	if err != nil {
		log.Fatal("Error setting heat", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)

	return sb
}

func getCurrentTemp() string {
	var currentTemp string
	resp, err := http.Get("http://10.0.0.24/aircon/get_control_info")

	if err != nil {
		log.Fatal("Error setting heat", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)

	items := strings.Split(sb, ",")

	for _, item := range items {
		if strings.Contains(item, "stemp") {
			currentTemp = strings.Split(item, "=")[1]
		}
	}

	return currentTemp
}

func getCurrentMode() string {
	var modeNum string
	resp, err := http.Get("http://10.0.0.24/aircon/get_control_info")

	if err != nil {
		log.Fatal("Error setting heat", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)

	items := strings.Split(sb, ",")

	for _, item := range items {
		if strings.Contains(item, "mode") {
			modeNum = strings.Split(item, "=")[1]
		}
	}

	return modeNum
}

func Toggle(state bool) {
	var mode string
	var stemp string
	var shum string
	var f_rate string
	var f_dir string

	onOrOff := map[bool]string{
		true:  "1",
		false: "0",
	}

	currentState := strings.Split(getCurrentState(), ",")

	for _, value := range currentState {
		fmt.Println(value)
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

	url := fmt.Sprintf("http://10.0.0.24/aircon/set_control_info?pow=%s&mode=%s&stemp=%s&shum=%s&f_rate=%s&f_dir=%s", onOrOff[state], mode, stemp, shum, f_rate, f_dir)

	fmt.Println(url)

	_, err := http.Get(url)

	if err != nil {
		log.Fatal("Error toggling aircon", err)
	}

}

func SetTemp(temp string) {
	_, err := http.Get(fmt.Sprintf("http://10.0.0.24/aircon/set_control_info?pow=1&mode=%s&stemp=%s&shum=0&f_rate=B&f_dir=3", getCurrentMode(), temp))

	if err != nil {
		log.Fatal("Error seting aircon temp", err)
	}
}

func SetHotOrCool(cool bool) {
	hotOrCold := map[bool]string{
		true:  "3",
		false: "4",
	}

	_, err := http.Get(fmt.Sprintf("http://10.0.0.24/aircon/set_control_info?pow=1&mode=%s&stemp=%s&shum=0&f_rate=B&f_dir=3", hotOrCold[cool], getCurrentTemp()))

	if err != nil {
		log.Fatal("Error toggling aircon", err)
	}
}

func SetFanRate(rate string) {
	rate = strings.Split(rate, "-")[1]

	fanSpeed := map[string]string{
		"night": "B",
		"1":     "3",
		"2":     "4",
		"3":     "5",
		"4":     "6",
		"5":     "7",
	}

	_, err := http.Get(fmt.Sprintf("http://10.0.0.24/aircon/set_control_info?pow=1&mode=%s&stemp=%s&shum=0&f_rate=%s&f_dir=3", getCurrentMode(), getCurrentTemp(), fanSpeed[rate]))

	if err != nil {
		log.Fatal("Error toggling aircon", err)
	}
}

type StatusStruct struct {
	temp     string
	mode     string
	fanSpeed string
	on       string
}

func Status() {
	var temp string
	var mode string
	var fanSpeed string
	var onOrOff string

	mapOnOrOfff := map[string]string{
		"0": "off",
		"1": "on",
	}

	mapModes := map[string]string{
		"3": "Cold",
		"4": "Hot",
	}

	mapFanSpeed := map[string]string{
		"B": "night",
		"3": "1",
		"4": "2",
		"5": "3",
		"6": "4",
		"7": "5",
	}

	resp, err := http.Get("http://10.0.0.24/aircon/get_control_info")

	if err != nil {
		log.Fatal("Error setting heat", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)

	items := strings.Split(sb, ",")

	for _, item := range items {
		if strings.Contains(item, "stemp") {
			temp = strings.Split(item, "=")[1]
		}

		if strings.Contains(item, "mode") {
			mode = mapModes[strings.Split(item, "=")[1]]
		}

		if strings.Contains(item, "f_rate") {
			fanSpeed = mapFanSpeed[strings.Split(item, "=")[1]]
		}

		if strings.Contains(item, "pow") {
			onOrOff = mapOnOrOfff[strings.Split(item, "=")[1]]
		}
	}

	fmt.Println(StatusStruct{temp, mode, fanSpeed, onOrOff})
}
