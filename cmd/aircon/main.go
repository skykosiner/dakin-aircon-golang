package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	a "github.com/skykosiner/aircon-control/pkg/aircon"
	"github.com/skykosiner/aircon-control/pkg/config"
	"github.com/skykosiner/aircon-control/pkg/utils"
	"github.com/spf13/cobra"
)

func printValidArgs(args []string, output io.Writer) {
	for _, arg := range args {
		fmt.Fprintln(output, arg)
	}
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "aircon",
		Short: "Dakin Aircon Controller",
	}

	config, err := config.GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	var verbose bool
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Have verbose outputs or not.")

	commands := []cobra.Command{
		{
			Use:   "status",
			Short: "Get the current statuso of the air con",
			Run: func(cmd *cobra.Command, args []string) {
				aircon := a.NewAircon(config.AirconIP, verbose)
				fmt.Println(aircon.StatusForUser())
			},
		},
		{
			Use:   "toggle",
			Short: "Toggle the power for the air con",
			Run: func(cmd *cobra.Command, args []string) {
				aircon := a.NewAircon(config.AirconIP, verbose)
				power := !utils.PowerToBool(aircon.Status.Power)
				aircon.SetStates(power, aircon.Status.Mode, aircon.Status.Temp, aircon.Status.Fan)
				aircon.SendRequest()
			},
		},
		{
			Use:   "mode",
			Short: "Toggle the pwore for the aircon",
			Run: func(cmd *cobra.Command, args []string) {
				aircon := a.NewAircon(config.AirconIP, verbose)
				validArgs := []string{"Cold", "Heat"}

				if len(args) == 0 {
					fmt.Fprintln(os.Stderr, "Please enter a valid arg.")
					printValidArgs(validArgs, os.Stderr)
					return
				}

				aircon.SetStates(utils.PowerToBool(aircon.Status.Power), args[0], aircon.Status.Temp, aircon.Status.Fan)
				aircon.SendRequest()
			},
		},
		{
			Use:   "fan",
			Short: "Set the fan mode.",
			Run: func(cmd *cobra.Command, args []string) {
				aircon := a.NewAircon(config.AirconIP, verbose)
				validArgs := []string{"Night", "1", "2", "3", "4", "5"}

				if len(args) == 0 {
					fmt.Fprintln(os.Stderr, "Please enter a valid arg.")
					printValidArgs(validArgs, os.Stderr)
					return
				}

				aircon.SetStates(utils.PowerToBool(aircon.Status.Power), aircon.Status.Mode, aircon.Status.Temp, args[0])
				aircon.SendRequest()
			},
		},
		{
			Use:   "temp",
			Short: "Set the temperature for the aircon",
			Run: func(cmd *cobra.Command, args []string) {
				aircon := a.NewAircon(config.AirconIP, verbose)
				validArgs := []string{
					"18",
					"18.5",
					"19",
					"19.5",
					"20",
					"20.5",
					"21",
					"21.5",
					"22",
					"22.5",
					"23",
					"23.5",
					"24",
					"24.5",
					"25",
					"25.5",
					"26",
					"26.5",
					"27",
					"27.5",
					"28",
					"28.5",
					"29",
					"29.5",
					"30",
					"30.5",
					"31",
				}

				if len(args) == 0 {
					fmt.Fprintln(os.Stderr, "Please enter a valid arg.")
					printValidArgs(validArgs, os.Stderr)
					return
				}

				aircon.SetStates(utils.PowerToBool(aircon.Status.Power), aircon.Status.Mode, args[0], aircon.Status.Fan)
				aircon.SendRequest()
			},
		},
	}

	for idx := range commands {
		rootCmd.AddCommand(&commands[idx])
	}

	if err := rootCmd.Execute(); err != nil {
		slog.Error("Command execution failed", "error", err)
		os.Exit(1)
	}
}
