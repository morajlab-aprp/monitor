package main

import (
	"fmt"
	// "os"
	"errors"
	"regexp"
	"strconv"
	"strings"
	// "os/exec"
)

type NumberT = int
type StringT = string
type BoolT = bool

type DumpString struct {
	ToNumber func(StringT) (NumberT, error)
	ToBool   func(StringT) (BoolT, error)
}

var dumpstring = DumpString{ToNumber: func(v StringT) (NumberT, error) {
	return strconv.Atoi(v)
}, ToBool: func(v StringT) (BoolT, error) {
	return strconv.ParseBool(v)
}}

func searchStrByRegex(regex, str string) []string {
	exp := regexp.MustCompile(regex)

	return exp.FindAllString(str, -1)
}

func getValueFromDump(key, dump string) (string, error) {
	matches := searchStrByRegex(fmt.Sprintf(`(?i)%s:\s*(\w|\d)+`, key), dump)

	if len(matches) > 1 {
		return "", errors.New("More than one match")
	}

	if len(matches) < 1 {
		return "", errors.New("no match")
	}

	tokens := strings.Split(matches[0], ":")
	value := strings.TrimSpace(tokens[1])

	return value, nil
}

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

	raw_value, err := getValueFromDump(`ac powered`, stdout)

	if err == nil {
		value, err := dumpstring.ToBool(raw_value)

		if err == nil {
			bs.ac_powered = value
		}
	}

	return &bs, nil
}

func main() {
	text, _ := dumpBattery()

	fmt.Printf("All matches : %+v\n", *text)
}
