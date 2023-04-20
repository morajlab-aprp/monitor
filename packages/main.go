package main

import (
	"fmt"
	// "os"
	"errors"
	"reflect"
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
