package enums

import "strings"

type Board int

const (
	UNKNOWN_BOARD Board = iota
	NOTICE
	FREE
	JOB
	PDS
	LECTURE
	BACHELOR
	SCHOLAR
)

var (
	boardMap = map[string]Board{
		"unknown":  UNKNOWN_BOARD,
		"notice":   NOTICE,
		"free":     FREE,
		"job":      JOB,
		"pds":      PDS,
		"lecture":  LECTURE,
		"bachelor": BACHELOR,
		"scholar":  SCHOLAR,
	}
)

func ParseBoard(str string) (Board, bool) {
	b, ok := boardMap[strings.ToLower(str)]
	return b, ok
}

func (b Board) String() string {
	return [...]string{"UNKNOWN_BOARD", "NOTICE", "FREE", "JOB", "PDS", "LECTURE", "BACHELOR", "SCHOLAR"}[b]
}

func (b Board) EnumIndex() int {
	return int(b)
}
