package main

import (
	"os"
	"strings"

	"github.com/skykosiner/aircon-control/pkg/aircon"
	"github.com/skykosiner/aircon-control/pkg/utils"
)

func main() {
	cmdArgs := os.Args[1:]

	switch cmdArgs[0] {
	case "toggle":
		if utils.GetOnState("10.0.0.24") {
			aircon.Toggle(false)
		} else {
			aircon.Toggle(true)
		}
	case "hot":
		aircon.SetHotOrCool(false)
	case "cold":
		aircon.SetHotOrCool(true)
	case "status":
		utils.CurrentStatus("10.0.0.24")
	case "conflict":
		aircon.FixConflict()
	default:
		if strings.Contains(cmdArgs[0], "fan") {
			aircon.SetFanRate(cmdArgs[0])
		} else {
			aircon.SetTemp(cmdArgs[0])
		}
	}
}
