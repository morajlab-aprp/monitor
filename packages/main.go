package main

import (
	"fmt"
	// "os"
	"regexp"
	"strings"
	// "os/exec"
)

type BatteryStats struct {
	ac_powered bool
}

func dumpBattery() (*BatteryStats, error) {
	var bs BatteryStats
	// cmd := exec.Command("dumpsys", "battery")
	// stdout, err := cmd.Output()

	// if err != nil {
	//   return "", err
	// }

	// return string(stdout), nil

	// TODO: Remove following line
	stdout := `Battery stat: Current Battery Service state:
  AC powered: true
  USB powered: false
  Wireless powered: false
  Max charging current: 0
  Max charging voltage: 0
  Charge counter: 0
  status: 2
  health: 2
  present: true
  level: 68
  scale: 100
  voltage: 0
  temperature: 0
  technology: Li-ion
`

	exp := regexp.MustCompile(`(?i)ac\spowered:\s*(false|true)\n`)
	matches := exp.FindAllString(stdout, -1)

	if len(matches) != 1 {
		// return custom error
	}

	return &bs, nil
}

func main() {
	text, _ := dumpBattery()

	fmt.Println("All matches : " + strings.Join(matches, " "))
}
