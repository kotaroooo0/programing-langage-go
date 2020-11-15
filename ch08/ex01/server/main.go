package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	port := flag.String("port", "8000", "port")
	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}

}

func handleConn(c net.Conn) {
	tz := os.Getenv("TZ")
	local, err := time.LoadLocation(tz)
	if err != nil {
		log.Fatal("error: invalid location")
	}

	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().In(local).Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
