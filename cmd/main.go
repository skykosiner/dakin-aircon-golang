package main

import (
	"flag"
	"fmt"

	a "github.com/skykosiner/aircon-control/pkg/aircon"
	"github.com/skykosiner/aircon-control/pkg/config"
	"github.com/skykosiner/aircon-control/pkg/utils"
)

func main() {
	// TODO: Figure out a better way to setup verbose mode
	// verbose := flag.Bool("v", false, "Verbose logging")
	verbose := true
	config, err := config.GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	aircon := a.NewAircon(config.AirconIP, verbose)
	power := flag.Bool("power", utils.PowerToBool(aircon.Status.Power), "Set the power state of the air con")
	mode := flag.String("mode", aircon.Status.Mode, "Air con hot or cold")
	temp := flag.String("temp", aircon.Status.Temp, "Set the temperature of the air con")
	fan := flag.String("fan", aircon.Status.Fan, "Set the fan mode of the air con")

	flag.Parse()
	aircon.SetStates(*power, *mode, *temp, *fan)
	aircon.SendRequest()
	fmt.Println(aircon.StatusForUser())
}
