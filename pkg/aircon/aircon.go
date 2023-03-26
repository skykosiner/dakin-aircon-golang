package aircon

import (
	"strings"

	"github.com/skykosiner/aircon-control/pkg/utils"
)

var mainAirconIp = utils.ReadConfig().MainIp
var conflictAirconOne = utils.ReadConfig().ConflictAirconOne
var conflictAirconTwo = utils.ReadConfig().ConflictAirconTwo

func Toggle(state bool) {
	onOrOff := map[bool]string{
		true:  "1",
		false: "0",
	}

	currentState := utils.CurrentStatus(mainAirconIp)
	utils.SendRequest(mainAirconIp, onOrOff[state],
		utils.MapValuesOfState(currentState.Mode),
		currentState.Temp,
		utils.MapValuesOfState(currentState.FanSpeed))
}

func SetTemp(temp string) {
	currentState := utils.CurrentStatus(mainAirconIp)
	utils.SendRequest(mainAirconIp, "1",
		utils.MapValuesOfState(currentState.Mode), temp,
		utils.MapValuesOfState(currentState.FanSpeed))
}

func SetHotOrCool(cool bool) {
	hotOrCold := map[bool]string{
		true:  "3",
		false: "4",
	}

	currentState := utils.CurrentStatus(mainAirconIp)
	utils.SendRequest(mainAirconIp, "1", hotOrCold[cool],
		currentState.Temp,
		utils.MapValuesOfState(currentState.FanSpeed))
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

	currentState := utils.CurrentStatus(mainAirconIp)
	utils.SendRequest(mainAirconIp, "1",
		utils.MapValuesOfState(currentState.Mode),
		currentState.Temp, fanSpeed[rate])
}

func FixConflict() {
	currentStateKitchen := utils.CurrentStatus(conflictAirconOne)
	currentStateRoom := utils.CurrentStatus(mainAirconIp)

	if currentStateKitchen.Power == "On" && currentStateKitchen.Mode != currentStateRoom.Mode {
		utils.SendRequest(conflictAirconOne, "0",
			utils.MapValuesOfState(currentStateKitchen.Mode),
			currentStateKitchen.Temp,
			utils.MapValuesOfState(currentStateKitchen.FanSpeed))

		// Set the other aircon downstairs to the same as the one that conflicts with mine
		utils.SendRequest(conflictAirconTwo, "1",
			utils.MapValuesOfState(currentStateKitchen.Mode),
			currentStateKitchen.Temp,
			utils.MapValuesOfState(currentStateKitchen.FanSpeed))
	}
}
