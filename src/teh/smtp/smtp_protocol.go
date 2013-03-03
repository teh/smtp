package smtp

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net"
)

const (
	VerbHELO              = 10
	VerbEHLO              = 11
	VerbRSET              = 12
	VerbRCPT              = 13
	VerbMAIL              = 14
	VerbQUIT              = 15
	VerbDATA              = 16
	VerbVRFY              = 17
	VerbSTARTTLS          = 18
	VerbAUTH_PLAIN        = 19
	VerbAUTH_PLAIN_DIRECT = 20
	VerbBASE64            = 21
)

// Default to 8BITMIME cause that's what everyone seems to use.
const (
	BodyType8BITMIME = 0
	BodyType7BIT     = 1
)

type Verb struct {
	Verb     int
	Data     []byte
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
	Hostname   string
	Cert       tls.Certificate
	remaining  []byte
	Parser     *Parser
	Recipients [][]byte
}

type Message struct {
	From       []byte
	Recipients [][]byte
	Body       []byte
}

type stateFunc func(c *Connection) stateFunc

func nextVerb(c *Connection) error {
	var data []byte = make([]byte, 1<<14)
	var err error
	var n int
	// clear data
	c.Parser.current.Data = []byte{}
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
		return err
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

func greet(c *Connection) stateFunc {
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
		return wait_for_tls
	case c.Parser.current.Verb == VerbEHLO:
		return ehlo_no_tls
	default:
		return wait_for_tls
	}
	panic("never reached")
}

func ioError(c *Connection) stateFunc {
	c.Close()
	return nil
}

func no_auth_envelope_empty(c *Connection) stateFunc {
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
		return no_auth_envelope_set
	case VerbVRFY:
		// http://cr.yp.to/smtp/vrfy.html
		fmt.Fprintf(c, "252 Send some mail to check.\r\n")
		return no_auth_envelope_empty
	case VerbAUTH_PLAIN:
		return auth_plain
	case VerbAUTH_PLAIN_DIRECT:
		return auth_plain_direct
	}
	log.Printf("503 bad command sequence [no_auth_envelope_empty]. Verb: %d", c.Parser.current.Verb)
	fmt.Fprintf(c, "503 bad command sequence\r\n")
	return no_auth_envelope_empty
}

func auth_ok_envelope_empty(c *Connection) stateFunc {
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
		return auth_ok_envelope_set
	case VerbVRFY:
		// http://cr.yp.to/smtp/vrfy.html
		fmt.Fprintf(c, "252 Send some mail to check.\r\n")
		return no_auth_envelope_empty
	}
	log.Printf("503 bad command sequence [auth_ok_envelope_empty]. Verb: %d", c.Parser.current.Verb)
	fmt.Fprintf(c, "503 bad command sequence\r\n")
	return no_auth_envelope_empty
}

func auth_plain_direct(c *Connection) stateFunc {
	user, _ := decodeUserAndMd5(c.Parser.current.Data)
	if user == nil {
		fmt.Fprintf(c, "501 syntax error\r\n")
		return ioError
	}
	// xxx check password
	fmt.Fprintf(c, "235 ok\r\n")
	return auth_ok_envelope_empty
}

// Decode user and md5-response from a base64 string.
func decodeUserAndMd5(data []byte) (user, md5_lower []byte) {
	n := base64.StdEncoding.DecodedLen(len(data))
	user_and_md5 := make([]byte, n)
	_, err := base64.StdEncoding.Decode(user_and_md5, data)
	if err != nil {
		return
	}

	split := bytes.SplitN(user_and_md5, []byte{0}, 2)
	if len(split) != 2 {
		return
	}
	user = split[0]
	md5_ := split[1]

	md5_lower = bytes.ToLower(md5_)
	return
}

func auth_plain(c *Connection) stateFunc {
	// important: keep the space after 334:
	fmt.Fprintf(c, "334 \r\n")
	err := nextVerb(c)
	if err != nil {
		log.Printf("Error on read: %s", err)
		return ioError
	}
	if c.Parser.current.Verb != VerbBASE64 {
		return ioError
	}

	_, _ = decodeUserAndMd5(c.Parser.current.Data)

	return auth_ok_envelope_empty
}


// In the no-auth case recipients can only be local where `local` is
// determined by the SAN DNSNames in the certificate.
func no_auth_envelope_set(c *Connection) stateFunc {
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
		return no_auth_envelope_set
	case VerbRCPT:
		// xxx no auth means only local recipients possible (check
		// against domain).
		fmt.Fprintf(c, "250 ok\r\n")
		return no_auth_envelope_set
	case VerbSTARTTLS:
		fmt.Fprintf(c, "220 go ahead\r\n")
		return starttls
	case VerbDATA:
		fmt.Fprintf(c, "354 ok\r\n")
		nextMessage(c)
		fmt.Fprintf(c, "250 ok\r\n")
		c.Recipients = nil // reset recipient list
		return no_auth_envelope_set
	}
	log.Printf("503 bad command sequence [no_auth_envelope_set]. Verb: %d", c.Parser.current.Verb)
	fmt.Fprintf(c, "503 bad command sequence\r\n")
	return no_auth_envelope_set
}

func auth_ok_envelope_set(c *Connection) stateFunc {
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
		return no_auth_envelope_set
	case VerbRCPT:
		fmt.Fprintf(c, "250 ok\r\n")
		return no_auth_envelope_set
	case VerbSTARTTLS:
		fmt.Fprintf(c, "220 go ahead\r\n")
		return starttls
	case VerbDATA:
		fmt.Fprintf(c, "354 ok\r\n")
		nextMessage(c)
		fmt.Fprintf(c, "250 ok\r\n")
		c.Recipients = nil // reset recipient list
		return no_auth_envelope_set
	}
	log.Printf("503 bad command sequence [no_auth_envelope_set]. Verb: %d", c.Parser.current.Verb)
	fmt.Fprintf(c, "503 bad command sequence\r\n")
	return no_auth_envelope_set
}

func ehlo_no_tls(c *Connection) stateFunc {
	// hardcoded, we don't care about configurability.
	extensions := []string{
		"250-8BITMIME",
		"250-STARTTLS",
		"250-ENHANCEDSTATUSCODES",
		"250 SIZE 33554432", // 32M
	}
	for _, e := range extensions {
		fmt.Fprintf(c, "%s\r\n", e)
	}
	return wait_for_tls
}

func wait_for_tls(c *Connection) stateFunc {
	err := nextVerb(c)
	if err != nil {
		log.Printf("Error on read: %s", err)
		return ioError
	}
	switch {
	case c.Parser.current.Verb == VerbQUIT:
		c.Close()
		return nil
	case c.Parser.current.Verb == VerbSTARTTLS:
		fmt.Fprintf(c, "220 go ahead\r\n")
		return starttls
	default:
		fmt.Fprintf(c, "530 Must STARTTLS first.\r\n")
		return wait_for_tls
	}
	panic("never reached")
}

func ehlo_with_tls(c *Connection) stateFunc {
	// hardcoded, we don't care about configurability.
	extensions := []string{
		"250-8BITMIME",
		"250-PIPELINING",
		"250-RSET",
		"250-NOOP",
		"250-VRFY",
		"250-STARTTLS",
		"250-ENHANCEDSTATUSCODES",
		"250-AUTH PLAIN",
		"250 SIZE 33554432", // 32M
	}
	for _, e := range extensions {
		fmt.Fprintf(c, "%s\r\n", e)
	}
	return no_auth_envelope_empty
}

func starttls(c *Connection) stateFunc {
	config := &tls.Config{
		Certificates:       []tls.Certificate{c.Cert},
		InsecureSkipVerify: true,
		//		ClientAuth: tls.RequireAnyClientCert,
	}
	// It's a TLS connection now, no going back.
	c.Conn = tls.Server(c.Conn, config)

	// Next thing after a TLS handshake is another EHLO.
	c.remaining = []byte("")
	err := nextVerb(c)
	if err != nil {
		log.Printf("Error on read: %s", err)
		return ioError
	}
	switch {
	case c.Parser.current.Verb == VerbHELO:
		fmt.Fprintf(c, "250 ok\r\n")
		return no_auth_envelope_empty
	case c.Parser.current.Verb == VerbEHLO:
		return ehlo_with_tls
	}
	log.Printf("Unexpected state after STARTTLS: %#v", c.Parser.current)
	return ioError
}

func Handle(c net.Conn, cert tls.Certificate) {
	conn := &Connection{
		Conn:     c,
		Hostname: "localhost",
		Parser:   NewParser(),
		Cert:     cert,
	}
	state := greet(conn)
	for {
		state = state(conn)
		if state == nil {
			return
		}
	}
}
