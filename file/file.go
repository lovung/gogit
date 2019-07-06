package file

import (
	"io/ioutil"
)

// ReadFile read whole file which have integers in every lines
func ReadFile(filename string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return data
}
