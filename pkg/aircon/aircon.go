package aircon

import (
	"strings"

	"github.com/skykosiner/aircon-control/pkg/utils"
)

func Toggle(state bool) {
	onOrOff := map[bool]string{
		true:  "1",
		false: "0",
	}

	currentState := utils.StoreCurrentState("10.0.0.24")
	utils.SendRequest("10.0.0.24", onOrOff[state], currentState.Mode, currentState.Stemp, currentState.F_rate)
}

func SetTemp(temp string) {
	currentState := utils.StoreCurrentState("10.0.0.24")
	utils.SendRequest("10.0.0.24", "1", currentState.Mode, temp, currentState.F_rate)
}

func SetHotOrCool(cool bool) {
	hotOrCold := map[bool]string{
		true:  "3",
		false: "4",
	}

	currentState := utils.StoreCurrentState("10.0.0.24")
	utils.SendRequest("10.0.0.24", "1", hotOrCold[cool], currentState.Stemp, currentState.F_rate)
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

	currentState := utils.StoreCurrentState("10.0.0.24")
	utils.SendRequest("10.0.0.24", "1", currentState.Mode, currentState.Stemp, fanSpeed[rate])
}

func FixConflict() {
	currentState := utils.StoreCurrentState("10.0.0.72")

	if currentState.Power {
		utils.SendRequest("10.0.0.72", "0", currentState.Mode, currentState.Stemp, currentState.F_rate)

		// Set the other aircon downstairs to the same as the one that conflicts with mine
		utils.SendRequest("10.0.0.54", "1", currentState.Mode, currentState.Stemp, currentState.F_rate)
	}
}
