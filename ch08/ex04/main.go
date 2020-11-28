package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	port := flag.String("port", "8000", "port")
	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		c, ok := conn.(*net.TCPConn)
		if !ok {
			log.Fatal("error: not TCPConn")
		}

		wg.Add(1)
		go func(conn *net.TCPConn) {
			defer wg.Done()
			input := bufio.NewScanner(conn)
			for input.Scan() {
				echo(conn, input.Text(), 1*time.Second)
			}
			conn.CloseWrite()
			defer conn.CloseRead()
		}(c)
	}

	// ?
	wg.Wait()
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
}
