package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1(args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2(args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

func timeEcho1() {
	start := time.Now()
	echo1(os.Args)
	fmt.Printf("%.8fs elapsed\n", time.Since(start).Seconds())
}

func timeEcho2() {
	start := time.Now()
	echo2(os.Args)
	fmt.Printf("%.8fs elapsed\n", time.Since(start).Seconds())
}

func timeEcho3() {
	start := time.Now()
	echo3(os.Args)
	fmt.Printf("%.8fs elapsed\n", time.Since(start).Seconds())
}

func main() {
	timeEcho1()
	timeEcho2()
	timeEcho3()
}
