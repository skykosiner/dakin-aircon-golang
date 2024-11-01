package utils

import (
	"fmt"
	"log"
	"net/http"
)

type StatusStruct struct {
	Temp     string
	Mode     string
	FanSpeed string
	Power    string
	F_dir    string
	Shum     string
}

type Conf struct {
	MainIp            string `json:"mainIp"`
	ConflictAirconOne string `json:"conflictAirconOne"`
	ConflictAirconTwo string `json:"conflictAirconTwo"`
}

func SendRequest(airconIp string, power string, mode string, temp string, fanRate string) {
	_, err := http.Get(fmt.Sprintf("http://%s/aircon/set_control_info?pow=%s&mode=%s&stemp=%s&shum=0&f_rate=%s&f_dir=3", airconIp, power, mode, temp, fanRate))

	if err != nil {
		log.Fatal("Error sending aircon request", err)
	}
}

func MapValuesOfState(item string) string {
	mapValues := map[string]string{
		// Modes
		"Cold": "3",
		"Heat": "4",
		// Power
		"On":  "1",
		"Off": "0",
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

func PowerToBool(power string) bool {
	return power == "On"
}
