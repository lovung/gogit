package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lovung/gogit/diff/lonp"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("./strdiff arg1 arg2")
	}
	diff := lonp.NewDiff(os.Args[1], os.Args[2])
	fmt.Printf("Editdistance: %d\n", diff.EditDistance())
	// fmt.Printf("LCS: %s\n", diff.LcsString())
	// fmt.Println("SES:")
	// diff.PrintSes()
}
