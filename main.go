package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	request(conn)

	// write response
	response(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			// request line
			m := strings.Fields(ln)[1]
			fmt.Println("URI: " + m)
		}
		if ln == "" {
			// headers are done
			break
		}

		i++
	}
}

func response(conn net.Conn) {
	body := "Hello World"

	fmt.Fprint(conn, "HTTP/1.1 200 OK \n")
	fmt.Fprintf(conn, "Content-Length: %d\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\n")
	fmt.Fprint(conn, "\n")
	fmt.Fprint(conn, body)
}
