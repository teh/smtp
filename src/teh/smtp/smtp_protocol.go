package smtp

import (
	"bytes"
	"net"
	"fmt"
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

type Connection struct {
	net.Conn
	Hostname string
	remaining []byte
	p *Parser
}

type stateFunc func(c Connection) stateFunc

func greet(c Connection) stateFunc {
	// Announce ESMTP, otherwise many clients won't EHLO.
	fmt.Fprintf(c, "220 %s ESMTP\r\n", c.Hostname)
	var data []byte
	for  {
		_, err := c.Read(data)
		if err != nil {
			return ioError
		}
		c.remaining, err = c.p.Feed(data)
		if err == Dangling {
			continue
		}
		if err == nil {
			switch {
			case c.p.current.Verb != VerbHELO:
				fmt.Fprintf(c, "250 ok\r\n")
				return normal
			case c.p.current.Verb != VerbEHLO:
				return ehlo
			default:
				return normal
				
			}
		}
		fmt.Fprintf(c, "500 could not parse HELO or EHLO.\r\n")	
	}
}

func ioError(c Connection) stateFunc {
	return ioError
}

func normal(c Connection) stateFunc {
	return normal
}

func ehlo(c Connection) stateFunc {
	extensions := []string{
		"250-8BITMIME",
		"250-PIPELINING",
		"250-RSET",
		"250-NOOP",
		"250-VRFY",
		"250 SIZE 33554432", // 32M
	}
	for _, e := range extensions {
		fmt.Fprintf(c, "%s\r\n", e)
	}
	return normal
}
