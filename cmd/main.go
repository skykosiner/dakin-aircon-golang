package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/skykosiner/aircon-control/pkg/aircon"
	helpinfo "github.com/skykosiner/aircon-control/pkg/helpInfo"
	"github.com/skykosiner/aircon-control/pkg/utils"
)

func main() {
	cmdArgs := os.Args[1:]

	switch cmdArgs[0] {
	case "toggle":
		if utils.CurrentStatus("10.0.0.24").Power == "On" {
			aircon.Toggle(false)
		} else {
			aircon.Toggle(true)
		}
	case "hot":
		aircon.SetHotOrCool(false)
	case "cold":
		aircon.SetHotOrCool(true)
	case "status":
		// Make sure connocted to the same wifi as the aircon
		networkName := exec.Command("iwgetid", "-r")
		stdOout, err := networkName.Output()

		if err != nil {
			log.Fatal("Error getting network name")
		}

		if strings.TrimSuffix(string(stdOout), "\n") == "The Kosiner's wifi" {
			fmt.Println(utils.CurrentStatus("10.0.0.24"))
		} else {
			fmt.Println("Not connected to correct wifi")
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
