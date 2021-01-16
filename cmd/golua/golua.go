package main

import (
	"io/ioutil"
	"os"

	"github.com/hilarryxu/golua/binchunk"
)

func main() {
	if len(os.Args) > 1 {
		data, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}

		binchunk.Undump(data)
	}
}
