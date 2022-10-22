package main

import (
	"fmt"
	"github.com/mrosset/scheme/pkg"
	"log"
)

func main() {
	version, err := scheme.EvalString("(version)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("guile version: %s\n", version)
}
