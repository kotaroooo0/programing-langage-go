package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs := os.Args[1:]
	if len(os.Args) == 1 {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		line := s.Text()
		inputs = strings.Split(line, " ")
	}

	for _, args := range inputs {
		t, err := strconv.ParseFloat(args, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "f: %v\n", err)
			os.Exit(1)
		}
		m := Meter(t)
		f := Feet(t)
		fmt.Printf("%s = %s , %s = %s\n", m, MeterToFeet(m), f, FeetToMeter(f))
	}
}

type Meter float64
type Feet float64

func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (f Feet) String() string  { return fmt.Sprintf("%gf", f) }

func MeterToFeet(m Meter) Feet { return Feet(3.2808 * m) }
func FeetToMeter(f Feet) Meter { return Meter(f / 3.2808) }
