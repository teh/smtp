package smtp

import "log"

const (
	VerbHELO = 10
	VerbEHLO = 11
	VerbRSET = 12
	VerbRCPT = 13
	VerbMAIL = 14
	VerbQUIT = 15
	VerbDATA = 16
	VerbVRFY = 17
)

type Verb struct {
	Verb     int
	Data       []byte
}

func emit(action int, data Verb) {
	log.Printf("[%d] S: %#v",
		action,
		string(data.Data),
	)
}

type Parser struct {
	cs int
	current Verb
	buffer [][]byte
	recording bool
}