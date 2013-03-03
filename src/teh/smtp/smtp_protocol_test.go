package smtp

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"net"
	smtp_client "net/smtp"
	"testing"
	"time"
)

func handle(c net.Conn, cert tls.Certificate) {
	conn := &Connection{
		Conn:     c,
		Hostname: "testhost",
		Parser:   NewParser(),
		Cert:     cert,
	}
	state := greet(conn)
	for {
		fmt.Printf("S %s %d, %#v\n", state, conn.Parser.current.Verb, string(conn.Parser.current.Data))
		state = state(conn)
		if state == nil {
			return
		}
	}
}

func TestEHLO(t *testing.T) {
	template := &x509.Certificate{
		Subject: pkix.Name{
			Country: []string{"uk"},
		},
		DNSNames:                    []string{"www.example.com"},
		NotBefore:                   time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
		NotAfter:                    time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC),
		SubjectKeyId:                []byte("6fcf9dfbd479ed82697fee719b9f8c610a11ff2a"),
		KeyUsage:                    x509.KeyUsageDataEncipherment,
		BasicConstraintsValid:       false,
		PermittedDNSDomainsCritical: false,
		PermittedDNSDomains:         []string{},
		SerialNumber:                big.NewInt(100),
		Version:                     0,
	}

	privkey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatal(err)
	}
	cert_data, err := x509.CreateCertificate(rand.Reader, template, template, &privkey.PublicKey, privkey)
	if err != nil {
		t.Fatal(err)
	}
	cert := tls.Certificate{Certificate: [][]byte{cert_data}}
	cert.PrivateKey = privkey

	s, c := net.Pipe()
	go handle(s, cert)
	client, err := smtp_client.NewClient(c, "localhost")
	if err != nil {
		t.Fatal(err)
	}
	err = client.Mail("someone@example.com")
	if err == nil {
		t.Errorf("Expected server to ask to STARTTLS first")
	}

	tls_client_config := tls.Config{
		Rand:               rand.Reader,
		ServerName:         "example.com",
		InsecureSkipVerify: true,
	}

	err = client.StartTLS(&tls_client_config)
	if err != nil {
		t.Error(err)
	}

	err = client.Mail("someone_new@example.com")
	if err != nil {
		t.Error(err)
	}
}
