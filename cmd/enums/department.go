package enums

import "strings"

type Department int

const (
	UNKNOWN_DEPARTMENT Department = iota
	ARCH
	CSE
	DORM
	MSE
	ACE
	IDE
	ITE
	MECHANICAL
	MECHATRONICS
	SCHOOL
	SIM
)

var (
	departmentMap = map[string]Department{
		"unknown":      UNKNOWN_DEPARTMENT,
		"arch":         ARCH,
		"cse":          CSE,
		"dorm":         DORM,
		"mse":          MSE,
		"ace":          ACE,
		"ide":          IDE,
		"ite":          ITE,
		"mechanical":   MECHANICAL,
		"mechatronics": MECHATRONICS,
		"school":       SCHOOL,
		"sim":          SIM,
	}
)

func ParseDepartment(str string) (Department, bool) {
	d, ok := departmentMap[strings.ToLower(str)]
	return d, ok
}

func (d Department) String() string {
	return [...]string{"UNKNOWN_DEPARTMENT", "ARCH", "CSE", "DORM", "MSE", "ACE", "IDE", "ITE", "MECHANICAL", "MECHATRONICS", "SCHOOL", "SIM"}[d]
}

func (d Department) EnumIndex() int {
	return int(d)
}
