// Exercise 1.3: Experiment to measure the difference in running time between
// our potentially inefficient versions and the one that uses strings.Join.
// (Section 1.6 illustrates part of the time package and Section 11.4 shows how
// to write benchmark tests for systematic performance evaluation.)

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// inefficient concatenation
	var s, sep string
	start := time.Now()
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println("concatenation:", time.Since(start))

	// strings.Join
	start = time.Now()
	_ = strings.Join(os.Args[1:], " ")
	fmt.Println("strings.Join:", time.Since(start))
}
