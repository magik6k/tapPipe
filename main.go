package main

import (
	"fmt"
	"github.com/songgao/water"
	"io"
	"os"
)

const BUFFERSIZE = 1522

func main() {
	ifce, err := water.NewTUN("")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "%v\n", ifce.Name())

	defer ifce.Close()

	go io.Copy(ifce, os.Stdin)
	io.Copy(os.Stdout, ifce)
}
