package main

// Test
import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/skykosiner/aircon-control/pkg/aircon"
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
	case "help":
		help := `Toggle air con on and off
aircon toggle

Set the temputure of the air con (any number between 18 - 30)
aircon 22

Set one of the fan modes
Fan modes available

night mode
level 1
level 2
level 3
level 4
level 5

aircon fan-night (or any number)

Get the current status of the air con
aircon status

Fix conflict between aircons that conflict with the one you want to control
aircon conflict
`

		fmt.Println(help)
	default:
		if strings.Contains(cmdArgs[0], "fan") {
			aircon.SetFanRate(cmdArgs[0])
		} else {
			aircon.SetTemp(cmdArgs[0])
		}
	}
}
