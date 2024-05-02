package main

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"

	http_connection "github.com/codecrafters-io/http-server-tester/internal/http/connection"
	http_request "github.com/codecrafters-io/http-server-tester/internal/http/request"
	http_response "github.com/codecrafters-io/http-server-tester/internal/http/response"
)

type Server struct {
	Listener net.Listener
}

func (s *Server) Listen() error {
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			return err
		}

		go s.handle(conn)
	}
}

var static = []byte("HTTP/1.1 200 OK\r\nContent-Length: 11\r\n\r\nhello world")

func (s *Server) handle(c net.Conn) {
	fmt.Printf("new conn!\n")
	buf := make([]byte, 1500)

	for {
		n, err := c.Read(buf)
		if err != nil {
			return
		}

		fmt.Println("read", n, "bytes")
		fmt.Println(string(buf[0:n]))
		_, _, err = http_request.Parse(buf[0:n])
		if err != nil {
			panic(err)
		}

		c.Write(static)
	}
}

// func main() {
// 	listener, err := net.Listen("tcp", ":8080")
// 	if err != nil {
// 		panic(err)
// 	}

// 	server := &Server{
// 		Listener: listener,
// 	}
// 	err = server.Listen()
// 	if err != nil {
// 		panic(err)
// 	}
// }

func main() {
	conn, err := http_connection.NewInstrumentedHttpConnection("localhost:5000", "client")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	req, _ := http.NewRequest("GET", "http://127.0.0.1:5000/text", bytes.NewBuffer([]byte("")))
	reqDump, _ := httputil.DumpRequestOut(req, true)

	conn.SendRequest(reqDump)

	data := make([]byte, 1024)
	n, err := conn.Conn.Read(data)
	if err != nil {
		panic(err)
	}
	response_bytes := (data[:n])
	r, _, err := http_response.Parse(response_bytes)
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}
