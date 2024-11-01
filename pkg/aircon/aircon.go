package aircon

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

type Status struct {
	Temp  string
	Mode  string
	Fan   string
	Power string
}

type Aircon struct {
	IP     string
	Status Status
}

func NewAircon(ip string, verbose bool) Aircon {
	state, err := currentState(ip)
	if err != nil {
		if verbose {
			slog.Error("Eror fetching aircon status", "error", err)
		}
		fmt.Println("No Status...")
		os.Exit(0)
	}

	return Aircon{
		IP:     ip,
		Status: state,
	}
}

func currentState(ip string) (Status, error) {
	var status Status

	url := fmt.Sprintf("http://%s/aircon/get_control_info", ip)
	resp, err := http.Get(url)
	if err != nil {
		return status, errors.New("Couldn't get the current air con status")
	}

	if resp.StatusCode != 200 {
		return status, errors.New("Status didn't return 200.")
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return status, errors.New("Couldn't read response from air con status")
	}

	states := strings.Split(string(bytes), ",")
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

	for _, state := range states {
		splitString := strings.Split(state, "=")
		item := splitString[0]
		value := splitString[1]

		switch item {
		case "stemp":
			status.Temp = value
		case "mode":
			status.Mode = modeMap[value]
		case "f_rate":
			status.Fan = fanMap[value]
		case "pow":
			status.Power = powerMap[value]
		}
	}

	return status, nil
}

func (a *Aircon) SetStates(power bool, mode string, temp string, fan string) {
	powerMap := map[bool]string{
		true:  "1",
		false: "0",
	}

	modeMap := map[string]string{
		"Cold": "3",
		"Heat": "4",
	}

	fanMap := map[string]string{
		"Night": "B",
		"1":     "3",
		"2":     "4",
		"3":     "5",
		"4":     "6",
		"5":     "7",
	}

	a.Status.Power = powerMap[power]
	a.Status.Mode = modeMap[mode]
	a.Status.Temp = temp
	a.Status.Fan = fanMap[fan]
}

func (a *Aircon) SendRequest() {
	_, err := http.Get(fmt.Sprintf("http://%s/aircon/set_control_info?pow=%s&mode=%s&stemp=%s&shum=0&f_rate=%s&f_dir=3", a.IP, a.Status.Power, a.Status.Mode, a.Status.Temp, a.Status.Fan))

	if err != nil {
		slog.Error("Erorr sending request to aircon", "error", err)
	}
}
