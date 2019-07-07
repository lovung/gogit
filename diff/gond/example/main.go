package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lovung/gogit/diff/gond"
	"github.com/lovung/gogit/file"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("./main file1 file2")
	}

	file1, _ := file.ReadLines(os.Args[1])
	file2, _ := file.ReadLines(os.Args[2])

	diff := gond.New(file1, file2)
	// diff.OnlyEd()
	diff.Compose()
	fmt.Printf("Editdistance: %d\n", diff.Editdistance())
	fmt.Printf("LCS: %s\n", diff.LcsString())
	fmt.Println("SES:")
	diff.PrintSes()
}
