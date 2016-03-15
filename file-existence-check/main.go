package main

import (
	"fmt"
	"os"
	"syscall"
)

func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func main() {
	fmt.Println("Hello", isExist("/usr/bin/env"), isExist("/usr/no-such-file"), isExist("/usr/bin"))

	var s syscall.Stat_t
	syscall.Stat("/usr/bin/env", &s)
	fmt.Println(s.Ctimespec.Sec)
	fmt.Println(fmt.Sprint(s.Ctimespec.Sec))
}
