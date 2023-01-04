package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type StatusStruct struct {
	Temp     string
	Mode     string
	FanSpeed string
	Power    string
	F_dir    string
	Shum     string
}

func SendRequest(airconIp string, power string, mode string, temp string, fanRate string) {
	_, err := http.Get(fmt.Sprintf("http://%s/aircon/set_control_info?pow=%s&mode=%s&stemp=%s&shum=0&f_rate=%s&f_dir=3", airconIp, power, mode, temp, fanRate))

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

func CurrentStatus(airconIp string) StatusStruct {
	var status StatusStruct

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
		parts := strings.Split(value, "=")

		switch parts[0] {
		case "stemp":
			status.Temp = parts[1]
		case "mode":
			status.Mode = modeMap[parts[1]]
		case "pow":
			status.Power = powerMap[parts[1]]
		case "f_rate":
			status.FanSpeed = fanMap[parts[1]]
		case "f_dir":
			status.F_dir = parts[1]
		case "shum":
			status.Shum = parts[1]
		}
	}

	return status
}

func MapValuesOfState(item string) string {
	mapValues := map[string]string{
		// Modes
		"Cold": "3",
		"Heat": "4",
		// Power
		"On":  "1",
		"Off": "2",
		// Fan speed
		"Night": "B",
		"1":     "3",
		"2":     "4",
		"3":     "5",
		"4":     "6",
		"5":     "7",
	}

	return mapValues[item]
}
