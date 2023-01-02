package main

import (
	"os"
	"strings"

	"github.com/skykosiner/aircon-control/pkg/aircon"
)

func main() {
	cmdArgs := os.Args[1:]

	switch cmdArgs[0] {
	case "on":
		aircon.Toggle(true)
	case "off":
		aircon.Toggle(false)
	case "hot":
		aircon.SetHotOrCool(false)
	case "cold":
		aircon.SetHotOrCool(true)
	case "status":
		aircon.Status()
	default:
		if strings.Contains(cmdArgs[0], "fan") {
			aircon.SetFanRate(cmdArgs[0])
		} else {
			aircon.SetTemp(cmdArgs[0])
		}
	}
}
