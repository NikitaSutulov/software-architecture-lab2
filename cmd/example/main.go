package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	lab2 "github.com/NikitaSutulov/software-architecture-lab2"
	"os"
)

func main() {
	var inConsole = flag.String("e", "", "input from console - compulsory input string")
	var inFileCaption = flag.String("f", "", "input from file - compulsory input file")
	var outFileCaption = flag.String("o", "", "output to file - compulsory output file")
	var err error
	var handler lab2.ComputeHandler

	flag.Parse()

	if err := checkFlags(inConsole, inFileCaption); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	handler = lab2.ComputeHandler{
		Input:  bytes.NewBufferString(*inConsole),
		Output: os.Stdout,
	}

	// Handle input from file
	if *inConsole == "" && *inFileCaption != "" {
		inFile, err := os.Open(*inFileCaption)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer inFile.Close()

		handler.Input = inFile
	}

	// Handle output to file
	if *outFileCaption != "" {
		outfile, _ := os.OpenFile(*outFileCaption, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
		handler.Output = outfile
		defer outfile.Close()
	}

	err = handler.Compute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Errors: %s", err)
	}
}

func checkFlags(e, f *string) error {
	if *e == "" && *f == "" {
		return errors.New("either -e or -f flag must be set")
	}
	if *e != "" && *f != "" {
		return errors.New("-e and -f flags cannot be used together")
	}
	return nil
}
