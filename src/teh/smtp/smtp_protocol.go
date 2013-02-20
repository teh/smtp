package smtp

import (
	"bytes"
	"net"
	"fmt"
	"log"
	"io"
	"crypto/tls"
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
	VerbSTARTTLS = 18
)

// Default to 8BITMIME cause that's what everyone seems to use.
const (
	BodyType8BITMIME = 0
	BodyType7BIT = 1
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
	Recipients [][]byte
}

type stateFunc func(c *Connection) stateFunc

func nextVerb(c *Connection) error {
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
			if err == io.EOF && len(c.remaining) == 0 {
				return err
			}
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

func nextMessage(c *Connection) error {
	var data []byte = make([]byte, 1<<14)
	var err error
	var n int
	mp := NewMessageParser()
	for {
		if len(c.remaining) != 0 {
			n = copy(data, c.remaining)
		} else {
			n, err = c.Read(data)
			if err == io.EOF && len(c.remaining) == 0 {
				return err
			}
			if err != nil && err != io.EOF {
				return err
			}
		}
		c.remaining, err = mp.Feed(data[:n])
		if err == Dangling {
			continue
		}
		if err == nil {
			return nil
		}
	}
	return nil
}

func Greet(c *Connection) stateFunc {
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
		return no_mail_yet
	case c.Parser.current.Verb == VerbEHLO:
		return ehlo
	default:
		return no_mail_yet
	}
	panic("never reached")
}

func ioError(c *Connection) stateFunc {
	c.Close()
	return nil
}

func no_mail_yet(c *Connection) stateFunc {
	err := nextVerb(c)
	if err != nil {
		log.Printf("Error on read: %s", err)
		return ioError
	}
	switch c.Parser.current.Verb {
	case VerbQUIT:
		c.Close()
		return nil
	case VerbMAIL:
		fmt.Fprintf(c, "250 ok\r\n")
		return envelope_set
	}
	log.Printf("503 bad command sequence\r\n")
	return no_mail_yet
}

func envelope_set(c *Connection) stateFunc {
	err := nextVerb(c)
	if err != nil {
		log.Printf("Error on read: %s", err)
		return ioError
	}
	switch c.Parser.current.Verb {
	case VerbQUIT:
		c.Close()
		return nil
	case VerbMAIL:
		fmt.Fprintf(c, "250 ok\r\n")
		c.Recipients = nil // reset recipient list
		return envelope_set
	case VerbRCPT:
		fmt.Fprintf(c, "250 ok\r\n")
		return envelope_set
	case VerbSTARTTLS:
		fmt.Fprintf(c, "220 go ahead\r\n")
		return starttls
	case VerbDATA:
		fmt.Fprintf(c, "354 ok\r\n")
		nextMessage(c)
		fmt.Fprintf(c, "250 ok\r\n")
		c.Recipients = nil // reset recipient list
		return envelope_set
	}
	log.Printf("503 bad command sequence\r\n")
	return envelope_set
}

func ehlo(c *Connection) stateFunc {
	// hardcoded, we don't care about configurability.
	extensions := []string{
		"250-8BITMIME",
		"250-PIPELINING",
		"250-RSET",
		"250-NOOP",
		"250-VRFY",
		"250-STARTTLS",
		"250-AUTH GSSAPI DIGEST-MD5 CRAM-MD5 PLAIN",
		"250 SIZE 33554432", // 32M
	}
	for _, e := range extensions {
		fmt.Fprintf(c, "%s\r\n", e)
	}
	return no_mail_yet
}

func starttls(c *Connection) stateFunc {
	err := nextVerb(c)
	if err != nil {
		log.Printf("Error on read: %s", err)
		return ioError
	}

	// go tls example: https://gist.github.com/spikebike/2233075
	// xxx key loading needs to go into servier init
    cert, err := tls.LoadX509KeyPair("certs/client.pem", "certs/client.key")
    if err != nil {
        log.Fatalf("server: loadkeys: %s", err)
    }
    config := &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	// It's a TLS connection now, no going back.
	c.Conn = tls.Server(c.Conn, config)

	// Start all over again after TLS handshake
	return Greet
}
