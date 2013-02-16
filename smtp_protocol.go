package smtp

import "log"

const (
	ActionHELO = iota
	ActionEHLO
	ActionRSET
	ActionRCPT
	ActionMAIL
)

type Action struct {
	Action     int
	Address    []byte
	ClientHost []byte
	Domain     []byte
}

func emit(action int, data Action) {
	log.Printf("[%d] host: `%s`, address: `%s`, domain: `%s`",
		action,
		string(data.ClientHost),
		string(data.Address),
		string(data.Domain),
	)
}
