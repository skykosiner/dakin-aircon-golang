package main

import (
	"os"
	"strings"

	"github.com/skykosiner/aircon-control/pkg/aircon"
)

func main() {
	cmdArgs := os.Args[1:]

	switch cmdArgs[0] {
	case "toggle":
		if aircon.GetOnState() {
			aircon.Toggle(false)
		} else {
			aircon.Toggle(true)
		}
	case "hot":
		aircon.SetHotOrCool(false)
	case "cold":
		aircon.SetHotOrCool(true)
	case "status":
		aircon.Status()
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
