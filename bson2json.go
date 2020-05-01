package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ma6174/bsonex"
)

func main() {
	parallel := flag.Int("p", 1, "parallel count")
	flag.Parse()
	var r io.Reader = os.Stdin
	var files []io.Reader
	for _, name := range flag.Args() {
		f, err := os.Open(name)
		if err != nil {
			log.Panicln(err)
		}
		files = append(files, f)
	}
	if len(files) > 0 {
		r = io.MultiReader(files...)
	}
	err := bsonex.NewDecoder(r).Do(*parallel, func(b bsonex.BSONEX) (err error) {
		_, err = fmt.Println(b)
		return err
	})
	if err != nil {
		log.Panicln(err)
	}
}
