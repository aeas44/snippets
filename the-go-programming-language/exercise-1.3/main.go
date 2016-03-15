// Experiment to measure the difference in running time between our potentially inefficient versions and the one that uses strings.Join. (Section 1.6 illustrates part of the time package, and Section 11.4 shows how to write benchmark tests for systematic performance evaluation.)

package main

import (
	"fmt"
	"strings"
	"time"
)

func myJoin(args []string) string {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}

	return s
}

func main() {
	stringsToJoin := []string{"to", "be", "or", "not", "to", "be"}
	repeatNumber := 30000000

	startTime := time.Now().Unix()
	for i := 0; i < repeatNumber; i++ {
		_ = strings.Join(stringsToJoin, " ")
	}

	middleTime := time.Now().Unix()

	for i := 0; i < repeatNumber; i++ {
		_ = myJoin(stringsToJoin)
	}

	finalTime := time.Now().Unix()
	fmt.Println("first benchmark (built-in): ", middleTime-startTime, "sec")
	fmt.Println("second benchmark (my function): ", finalTime-middleTime, "sec")
}
