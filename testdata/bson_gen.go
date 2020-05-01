package main

import (
	"fmt"
	"log"

	"github.com/globalsign/mgo/bson"
)

func main() {
	b, err := bson.Marshal(map[string]interface{}{
		"int":    int(123),
		"float":  float64(4.5),
		"array":  []string{"a", "b"},
		"doc":    map[string]int{"c": 1},
		"string": "def",
	})
	if err != nil {
		log.Panicln(err)
	}
	fmt.Print(string(b))
	b, err = bson.Marshal([]string{"a", "b"})
	if err != nil {
		log.Panicln(err)
	}
	fmt.Print(string(b))
}
