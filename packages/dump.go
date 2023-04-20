package main

type DumpValue string
type Dump struct {
	Stdout string
}

func (dump Dump) GetValue(key string) (DumpValue, error) {
	matches := SearchStrByRegex(fmt.Sprintf(`(?i)%s:\s*(\w|\d)+`, key), string(dump))

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

func (dv DumpValue) ToNumber() (NumberT, error) {
	return strconv.Atoi(string(dv))
}

func (dv DumpValue) ToBool() (BoolT, error) {
	return strconv.ParseBool(string(dv))
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
