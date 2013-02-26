/* On startup we start listening. On SIGINT and SIGTERM we stop
listening immediately but shut down with the foll

*/
package main

import (
	"teh/smtp"
	"net"
	"flag"
	"log"
	"os/signal"
	"os"
	"time"
	"syscall"
	"fmt"
	"crypto/tls"
	"crypto/x509"
)

type Config struct {
	ListenAddress *string
	Cert tls.Certificate
}

type Server struct {
	stop bool
}

func (server *Server) listen(listener net.Listener, newConnections chan net.Conn) {
	for {
		c, err := listener.Accept()

		if err == nil {
			log.Printf("New connection from `%s`", c.RemoteAddr())
			newConnections <- c
			continue
		}

		if ne, ok := err.(net.Error); ok && ne.Temporary() {
			log.Printf("Temporary error on accept: %s. Ignoring.", err)
			continue
		}
		
		if server.stop {
			log.Printf("Shutting down listener because of server stop.")
 		} else {
			log.Printf("Fatal error on accept: %s. Shutting down listener.", err)
		}
		return
	}
}

func configure() *Config {
	pem := flag.String("tls_pem", "", "(required) Path to the TLS public key (PEM).")
	key := flag.String("tls_key", "", "(required) Path to the TLS private key.")
	listen_address := flag.String("listen_address", "127.0.0.1:2000", "Listen address, e.g. :25")

	flag.Parse()

	if *pem == "" {
		log.Fatalf("Argument tls_pem is required.")
	}
	if *key == "" {
		log.Fatalf("Argument tls_key is required.")
	}
	cert, err := tls.LoadX509KeyPair(*pem, *key)
	if err != nil {
		log.Fatalf("Could not load certificate: %s", err)
	}

	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		log.Fatalf("Could not parse certificate %s", err)
	}

	return &Config{
		ListenAddress: listen_address,
		Cert: cert,
	}
}

func main() {
	config := configure()

	listener, err := net.Listen("tcp", *config.ListenAddress)

	if err != nil {
		log.Fatalf("Could not listen on `%s`: `%s`", *config.ListenAddress, err)
	} else {
		log.Printf("Listening on `%s` for domains: %v", *config.ListenAddress, config.Cert.Leaf.DNSNames)
	}
	newConnections := make(chan net.Conn)
	signals := make(chan os.Signal, 2)

	server := Server{}

	go server.listen(listener, newConnections)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	for !server.stop {
		select {
		case s := <-signals:
			log.Printf("Recevied signal `%s`, shutting down.", s)
			server.stop = true
			listener.Close()
		case newConn := <-newConnections:
			go smtp.Handle(newConn, config.Cert)
		}
	}
	time.Sleep(time.Second)
}
