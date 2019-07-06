package main

import (
	"fmt"

	"github.com/lovung/gogit/file"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func a() {
	dmp := diffmatchpatch.New()

	origin := "test/origin.go"
	edited := file.ReadFile("test/edited.go")

	diffs := dmp.DiffMain(string(origin), string(edited), true)

	fmt.Println(dmp.DiffPrettyText(diffs))
}
