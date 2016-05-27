package main

import (
	"fmt"
	"github.com/songgao/water"
	"io"
	"os"
)

const BUFFERSIZE = 1522

func main() {
	ifce, err := water.NewTAP("")
	tapstd := make([]byte, BUFFERSIZE)
	stdtap := make([]byte, BUFFERSIZE)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "%v\n", ifce.Name())

	defer ifce.Close()

	go func() {
		n, err := os.Stdin.Read(stdtap)
		if err != nil && n == 0 {
			ifce.Close()
			if err == io.EOF {
				os.Exit(0)
			} else {
				os.Exit(1)
			}

		}

		ifce.Write(stdtap)
	}()

	for {
		n, err := ifce.Read(tapstd)
		if err != nil && n == 0 {
			return
		}

		os.Stdout.Write(tapstd)
	}
}
