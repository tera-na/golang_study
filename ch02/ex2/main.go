// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/adonovan/gopl.io/ch2/tempconv"
)

func main() {
	value = os.Args[1:]
	if len(value) == 0 {
		input := bufio.NewScanner(os.Stdin)
		convCtoAny(input)
	} else {
		for _, arg := range os.Args[1:] {
			convCtoAny(arg)
		}
	}
}

func convCtoAny( arg string ) {
	t, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c))
}

//!-
