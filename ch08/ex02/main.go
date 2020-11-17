package main

import (
	"log"

	"github.com/kotaroooo0/programing-language-go/ch08/ex02/orenoftp"
)

func main() {
	s := orenoftp.Server{
		Addr: ":2121",
	}
	log.Fatal(s.ListenAndServe())
}
