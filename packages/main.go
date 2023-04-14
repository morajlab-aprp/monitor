package main

import (
	"fmt"
	// "os"
	"errors"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	// "os/exec"
)

type NumberT = int
type StringT = string
type BoolT = bool

const NumberTName = "int"
const StringTName = "string"
const BoolTName = "bool"

type DumpValue string
type Dump string

func (dv DumpValue) ToNumber() (NumberT, error) {
	return strconv.Atoi(string(dv))
}

func (dv DumpValue) ToBool() (BoolT, error) {
	return strconv.ParseBool(string(dv))
}

type Services struct {
	Battery string
}

var services = Services{Battery: "battery"}

func searchStrByRegex(regex, str string) []string {
	exp := regexp.MustCompile(regex)

	return exp.FindAllString(str, -1)
}

func (dump Dump) getValue(key string) (DumpValue, error) {
	matches := searchStrByRegex(fmt.Sprintf(`(?i)%s:\s*(\w|\d)+`, key), string(dump))

	if len(matches) > 1 {
		return DumpValue(""), errors.New("More than one match")
	}

	if len(matches) < 1 {
		return DumpValue(""), errors.New("no match")
	}

	tokens := strings.Split(matches[0], ":")
	value := strings.TrimSpace(tokens[1])

	return DumpValue(value), nil
}

func dumpSystem(service string) (Dump, error) {
	// cmd := exec.Command("dumpsys", service)
	// stdout, err := cmd.Output()

	// return Dump(stdout), err

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

	return Dump(stdout), nil
}

type DumpField struct {
	Key   string
	Type  string
	Value DumpValue
}

func (field *DumpField) Set(dump Dump) error {
	dumpvalue, err := dump.getValue(field.Key)

	if err != nil {
		return err
	}

	field.Value = dumpvalue

	return nil
}

func (field *DumpField) Get() (any, error) {
	if field.Type == BoolTName {
		return field.Value.ToBool()
	}

	if field.Type == StringTName {
		return field.Value, nil
	}

	return nil, fmt.Errorf("Type %s is invalid", field.Type)
}

type BatteryStats struct {
	Ac_powered DumpField
	Technology DumpField
}

// TODO: Complete here
func (bs BatteryStats) Initialize(dump Dump) error {
	v := reflect.ValueOf(bs)
	typeOfbs := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Println(v.Field(i).Interface().(DumpField))
		fmt.Printf("Field: %s\tValue: %v\n", typeOfbs.Field(i).Name, v.Field(i).Interface().(DumpField).Value)
	}

	return nil
}

func dumpBattery() (*BatteryStats, error) {
	bs := &BatteryStats{Ac_powered: DumpField{Key: "ac powered", Type: BoolTName}, Technology: DumpField{Key: "technology", Type: StringTName}}
	dump, err := dumpSystem(services.Battery)

	if err != nil {
		return bs, nil
	}

	err = bs.Initialize(dump)

	if err != nil {
		return bs, nil
	}

	return bs, nil
}

func main() {
	text, _ := dumpBattery()

	// val, _ := text.Ac_powered.Get()
	fmt.Printf("All matches : %+v\n", text)
}
