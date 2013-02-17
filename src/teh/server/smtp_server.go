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
)

type Server struct {
	stop bool
}

func handle(c net.Conn) {
	logger := log.New(os.Stdout, fmt.Sprintf("[%s] ", c.RemoteAddr()) , log.Ldate + log.Ltime)
	logger.Printf("Handler started.")
	conn := smtp.Connection{
		Conn: c,
		Hostname: "testhost",
		Parser: smtp.NewParser(),
	}
	state := smtp.Greet(conn)
	for {
		state = state(conn)
		if state == nil {
			return
		}
	}
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

func main() {
	listen_address := flag.String("listen_address", "127.0.0.1:2000", "Listen address, e.g. :25")
	listener, err := net.Listen("tcp", *listen_address)

	if err != nil {
		log.Fatalf("Could not listen on `%s`: `%s`", *listen_address, err)
	} else {
		log.Printf("Listening on `%s`", *listen_address)
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
			go handle(newConn)
		}
	}
	time.Sleep(time.Second)
}
