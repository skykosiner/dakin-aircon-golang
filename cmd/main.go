package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/skykosiner/aircon-control/pkg/aircon"
	helpinfo "github.com/skykosiner/aircon-control/pkg/helpInfo"
	"github.com/skykosiner/aircon-control/pkg/utils"
)

func main() {
	cmdArgs := os.Args[1:]
	mainAirconIp := utils.ReadConfig().MainIp

	switch cmdArgs[0] {
	case "toggle":
		if utils.CurrentStatus(mainAirconIp).Power == "On" {
			aircon.Toggle(false)
		} else {
			aircon.Toggle(true)
		}
	case "hot":
		aircon.SetHotOrCool(false)
	case "cold":
		aircon.SetHotOrCool(true)
	case "status":
		curr := utils.CurrentStatus(mainAirconIp)
		if curr == (utils.StatusStruct{}) {
			fmt.Println("No Status")
		} else {
			fmt.Printf("%s %s %s %s\n", curr.Temp, curr.Mode, curr.FanSpeed, curr.Power)
		}
	case "conflict":
		aircon.FixConflict()
	case "setupHelp":
		helpinfo.MoveHelpFile()
	case "help":
		homePath := os.Getenv("HOME")
		bytes, err := os.ReadFile(fmt.Sprintf("%s/.local/airconhelp.txt", homePath))

		if err != nil {
			log.Fatal("Error reading help file", err)
		}

		fmt.Println(string(bytes))
	default:
		if strings.Contains(cmdArgs[0], "fan") {
			aircon.SetFanRate(cmdArgs[0])
		} else {
			aircon.SetTemp(cmdArgs[0])
		}
	}
}
