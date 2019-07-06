package main

import (
	"fmt"

	"github.com/lovung/gogit/file"
	"github.com/sergi/go-diff/diffmatchpatch"
	// "strings"
	// "github.com/sergi/go-diff/diffmatchpatch"
)

const (
	text1 = "Actually source code in C can of course be found in e.g. the GNU version of diff or the one from Version 7 Unix, also available online. The core of diff algorithms seeks to compare two sequences and to discover how the first can be transformed into the second by a sequence of operations using the primitives delete-subsequence, and insert-subseqence. If a delete and an insert coincide on the same range then it can be labeled as a change-subsequence. The user doesn't always want to be informed as to which subsequences remain identical, but this information is nonetheless typically computed as a by-product of the basic algorithm."
	text2 = "The core of diff algorithms seeks to compare two sequences and to discover how the first can be transformed into the second by a sequence of operations using the primitives delete-subsequence, and insert-subseqence. If a delete and an insert coincide on the same range then it can be labeled as a change-subsequence. Once identical elements are found, they are clumped together into runs which were adjacent in the pre-sorted sequences. The user doesn't always want to be informed as to which subsequences remain identical, but this information is nonetheless typically computed as a by-product of the basic algorithm."
)

func main() {
	dmp := diffmatchpatch.New()

	origin := file.ReadFile("test/origin.go")
	edited := file.ReadFile("test/edited.go")

	diffs := dmp.DiffMain(string(edited), string(origin), true)

	fmt.Println(dmp.DiffPrettyText(diffs))
}
