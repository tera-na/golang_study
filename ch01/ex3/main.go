package main

import (
	"fmt"
	"strings"
	"time"
)

func echo1(args ...string){
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2(args ...string){
	fmt.Println(strings.Join(args[0:], " "))
}

func main() {

	start := time.Now()
	for i:=0; i <10000; i+=1{
		echo1("ssss","sss")
	}
	sec1 := time.Since(start).Seconds()

	start = time.Now()
	for i:=0; i <10000; i+=1{
		echo2("ssss","sss")
	}
	sec2 := time.Since(start).Seconds()
	
	fmt.Printf("%.2fs", sec2-sec1)
}

