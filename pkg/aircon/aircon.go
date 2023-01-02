package aircon

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func getCurrentTemp() string {
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
		if strings.Contains(item, "stemp") {
			modeNum = strings.Split(item, "=")[1]
		}
	}

	return modeNum
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
	onOrOff := map[bool]string{
		true:  "1",
		false: "0",
	}

	_, err := http.Get(fmt.Sprintf("http://10.0.0.24/aircon/set_control_info?pow=%s&mode=4&stemp=26&shum=0&f_rate=B&f_dir=3", onOrOff[state]))

	if err != nil {
		log.Fatal("Error toggling aircon", err)
	}
}

func SetTemp(temp string) {
	_, err := http.Get(fmt.Sprintf("http://10.0.0.24/aircon/set_control_info?pow=1&mode=%s&stemp=%s&shum=0&f_rate=B&f_dir=3", getCurrentMode(), temp))

	if err != nil {
		log.Fatal("Error toggling aircon", err)
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
