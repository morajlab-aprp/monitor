package main

import (
	"fmt"
	"regexp"
)

type Regex struct {
	Value string
}

func (reg *Regex) SearchStr(str string) []string {
	exp := regexp.MustCompile(reg.Value)

	return exp.FindAllString(str, -1)
}

func CreateRegex(pattern, flag string) Regex {
	return &Regex{Value: fmt.Sprintf(`(?%s)%s`, flag, pattern)}
}
