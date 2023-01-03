package aircon

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/skykosiner/aircon-control/pkg/utils"
)

type StatusStruct struct {
	temp     string
	mode     string
	fanSpeed string
	on       string
}

func Toggle(state bool) {
	onOrOff := map[bool]string{
		true:  "1",
		false: "0",
	}

	currentState := utils.StoreCurrentState("10.0.0.24")

	url := fmt.Sprintf("http://10.0.0.24/aircon/set_control_info?pow=%s&mode=%s&stemp=%s&shum=%s&f_rate=%s&f_dir=%s", onOrOff[state], currentState.Mode, currentState.Stemp, currentState.Shum, currentState.F_rate, currentState.F_dir)
	_, err := http.Get(url)

	if err != nil {
		log.Fatal("Error toggling aircon", err)
	}

}

func SetTemp(temp string) {
	_, err := http.Get(fmt.Sprintf("http://10.0.0.24/aircon/set_control_info?pow=1&mode=%s&stemp=%s&shum=0&f_rate=B&f_dir=3", utils.GetCurrentMode("10.0.0.24"), temp))

	if err != nil {
		log.Fatal("Error seting aircon temp", err)
	}
}

func SetHotOrCool(cool bool) {
	hotOrCold := map[bool]string{
		true:  "3",
		false: "4",
	}

	_, err := http.Get(fmt.Sprintf("http://10.0.0.24/aircon/set_control_info?pow=1&mode=%s&stemp=%s&shum=0&f_rate=B&f_dir=3", hotOrCold[cool], utils.GetCurrentTemp("10.0.0.24")))

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

	_, err := http.Get(fmt.Sprintf("http://10.0.0.24/aircon/set_control_info?pow=1&mode=%s&stemp=%s&shum=0&f_rate=%s&f_dir=3", utils.GetCurrentMode("10.0.0.24"), utils.GetCurrentTemp("10.0.0.24"), fanSpeed[rate]))

	if err != nil {
		log.Fatal("Error toggling aircon", err)
	}
}

func Status() {
	var temp string
	var mode string
	var fanSpeed string
	var onOrOff string

	mapOnOrOfff := map[string]string{
		"0": "Off",
		"1": "On",
	}

	mapModes := map[string]string{
		"3": "Cold",
		"4": "Hot",
	}

	mapFanSpeed := map[string]string{
		"B": "Night",
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

func FixConflict() {
	// Turn off kitchen aircon
	currentState := utils.StoreCurrentState("10.0.0.24")

	url := fmt.Sprintf("http://10.0.0.72/aircon/set_control_info?pow=0&mode=%s&stemp=%s&shum=%s&f_rate=%s&f_dir=%s", currentState.Mode, currentState.Stemp, currentState.Shum, currentState.F_rate, currentState.F_dir)
	_, err := http.Get(url)

	if err != nil {
		log.Fatal("Error toggling aircon", err)
	}
}
