package main

import (
	"fmt"

	"github.com/lovung/gogit/diff/gonp"
	"github.com/lovung/gogit/file"
)

func main() {
	// if len(os.Args) < 3 {
	// 	log.Fatal("./strdiff arg1 arg2")
	// }
	origin := file.ReadFile("test/origin.go")
	edited := file.ReadFile("test/edited.go")
	diff := gonp.New(string(edited), string(origin))
	// diff.OnlyEd()
	diff.Compose()
	fmt.Printf("Editdistance: %d\n", diff.Editdistance())
	fmt.Printf("LCS: %s\n", diff.LcsString())
	fmt.Println("SES:")
	diff.PrintSes()
}
