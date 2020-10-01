package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	hash := flag.String("hash", "sha256", "sha256 or sha384 or sha512")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	bytes := scanner.Bytes()

	switch *hash {
	case "sha256":
		sum := sha256.Sum256(bytes)
		fmt.Println(hex.EncodeToString(sum[:]))
	case "sha384":
		sum := sha512.Sum384(bytes)
		fmt.Println(hex.EncodeToString(sum[:]))
	case "sha512":
		sum := sha512.Sum512(bytes)
		fmt.Println(hex.EncodeToString(sum[:]))
	}
}
