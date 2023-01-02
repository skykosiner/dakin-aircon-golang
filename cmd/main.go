package main

import (
	"os"

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
	default:
		aircon.SetTemp(cmdArgs[0])
	}
}
