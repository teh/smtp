package smtp

import (
	"bytes"
	"net"
	"fmt"
	"log"
	"io"
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

const (
	BodyType8BITMIME = 1
	BodyType7BIT = 2
)

type Verb struct {
	Verb int
	Data []byte
	BodyType int
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
	Parser *Parser
}

type stateFunc func(c Connection) stateFunc

func nextVerb(c Connection) error {
	var data []byte = make([]byte, 1<<14)
	var err error
	var n int
	for {
		// remaining always gets parsed first before reading new data.
		// This is important not just for ordering but also because
		// remaining is a sub-slice of data (copy handles overlapping
		// regions correctly).
		if len(c.remaining) != 0 {
			n = copy(data, c.remaining)
		} else {
			n, err = c.Read(data)
			if err != nil && err != io.EOF {
				return err
			}
		}
		c.remaining, err = c.Parser.Feed(data[:n])
		if err == Dangling {
			continue
		}
		if err == nil {
			return nil
		}
	}
	return nil
}

func Greet(c Connection) stateFunc {
	// Announce ESMTP, otherwise many clients won't EHLO.
	fmt.Fprintf(c, "220 %s ESMTP\r\n", c.Hostname)
	err := nextVerb(c)
	if err != nil {
		log.Printf("Error on read: %s", err)
		return ioError
	}
	switch {
	case c.Parser.current.Verb == VerbHELO:
		fmt.Fprintf(c, "250 ok\r\n")
		return normal
	case c.Parser.current.Verb == VerbEHLO:
		return ehlo
	default:
		return normal
	}
	panic("never reached")
}

func ioError(c Connection) stateFunc {
	return nil
}

func normal(c Connection) stateFunc {
	for {
		err := nextVerb(c)
		if err != nil {
			log.Printf("Error on read: %s", err)
			return ioError
		}
		switch c.Parser.current.Verb {
		case VerbQUIT:
			c.Close()
			return nil
		default:
		}
	}
	return normal
}

func ehlo(c Connection) stateFunc {
	// hardcoded, we don't care about configurability.
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
