package smtp

import (
	"bytes"
)

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
	Verb int
	Data []byte
}

type Parser struct {
	cs      int
	current Verb
	buffer  *bytes.Buffer
}

type MessageParser struct {
	cs     int
	buffer *bytes.Buffer
}
