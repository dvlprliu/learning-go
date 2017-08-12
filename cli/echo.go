// Echo prints its command-line arguments.
package main

import (
	"strings"
	"fmt"
	"os"
)

func main() {
	s := echo1()
	fmt.Println(s)
	s1 := echo2() 
	fmt.Println(s1)
	s2 := echo3()
	fmt.Println(s2)
	exercise1 := exerciseEcho()
	fmt.Println("exercise1.1: ", exercise1)
	exerciseEcho2()
}

func echo1() string {
	var s, sep string
	for i:=1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func echo2() string {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func echo3() string {
	return strings.Join(os.Args[1:], " ")
}

// print the name of command tha invoked it
func exerciseEcho() string {
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	return s
}

func exerciseEcho2() {
	for idx, arg := range os.Args {
		fmt.Println("idx: ", idx, " arg: ", arg)
	}
}
