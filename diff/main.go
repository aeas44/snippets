package main
import (
	"github.com/sergi/go-diff/diffmatchpatch"
	"fmt"
)

func lineDiff(src1, src2 string) []diffmatchpatch.Diff {
	dmp := diffmatchpatch.New()

	a, b, c := dmp.DiffLinesToChars(src1, src2)
	diffs := dmp.DiffMain(a, b, false)
	result := dmp.DiffCharsToLines(diffs, c)
	fmt.Println(result)

	return result
}

func main() {
	var a = `
	abc
	def
	ghi
	`
	var b = `
	abc
	go-
	ghi
	`
	lineDiff(a, b)
}
