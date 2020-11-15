package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	newYorkPort := flag.String("NewYork", "8010", "NewYork Port")
	tokyoPort := flag.String("Tokyo", "8020", "Tokyo Port")
	londonPort := flag.String("London", "8030", "London Port")
	flag.Parse()

	go displayTime("localhost:" + *newYorkPort)
	go displayTime("localhost:" + *tokyoPort)
	displayTime("localhost:" + *londonPort)
}

func displayTime(host string) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
