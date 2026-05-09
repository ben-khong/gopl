// Exercise 1.2: Modify the echo program to print the index and value of each of its arguments, one per line.
package main

import (
	"fmt"
	"os"
)

func main() {
	s := ""
	for i, arg := range os.Args[1:] {
		fmt.Println(i+1, arg)
	}
	fmt.Println(s)

}
