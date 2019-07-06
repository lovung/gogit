package main

import (
	"fmt"

	"github.com/lovung/gogit/diff/gond"
	dmp "github.com/sergi/go-diff/diffmatchpatch"
)

func main() {
	// if len(os.Args) < 3 {
	// 	log.Fatal("./strdiff arg1 arg2")
	// }
	var constitution []string
	constitution = append(constitution, "We the People of the United States, in Order to form a more perfect Union,")
	constitution = append(constitution, "establish Justice, insure domestic Tranquility, provide for the common defence,")
	constitution = append(constitution, "promote the general Welfare, and secure the Blessings of Liberty to ourselves")
	constitution = append(constitution, "and our Posterity, do ordain and establish this Constitution for the United")
	constitution = append(constitution, "States of America.")

	var got []string
	got = append(got, "TEST")
	got = append(got, "We the People of the United States, in Order to form a more perfect Union,")
	got = append(got, "establish Justice, insure domestic Tranquility, provide for the common defence,")
	got = append(got, "and secure the Blessings of Liberty to ourselves")
	got = append(got, "and our Posterity, do ordain and establish this Constitution for the United")
	got = append(got, "States of America.")

	diff := gond.New(constitution, got)
	// diff.OnlyEd()
	diff.Compose()
	fmt.Printf("Editdistance: %d\n", diff.Editdistance())
	fmt.Printf("LCS: %s\n", diff.LcsString())
	fmt.Println("SES:")
	diff.PrintSes()

	dmp := diffmatchpatch.New()
	diffs := dmp.
	fmt.Println(dmp.DiffPrettyText(diffs))
}
