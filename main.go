package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	var s,sep string
	a := "AAA"
	var b = "BBB"
	var c string
	c = "CCC"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	 for i := 0; i<len(os.Args); i++{
	 	s += sep + os.Args[i]
		 fmt.Printf("os args params key %d value %v \n", i, os.Args[i])
	 	sep = " # "
	 }
	fmt.Println(strings.Join(os.Args, "I"))
	fmt.Println(s)
}