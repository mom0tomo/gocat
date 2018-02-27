package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var line string

func usage() {
	fmt.Fprintf(os.Stderr, "usage: gocat [file...]\n")
	os.Exit(2)
}

func main() {
	log.SetPrefix("gocat: ")
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parsegi()

	if flag.NArg() == 0 {
		read(os.Stdin)
	} else {
		for _, arg := range flag.Args() {
			f, err := os.Open(arg)
			if err != nil {
				log.Print(err)
				continue
			}
			read(f)
			f.Close()
		}
	}
	print(os.Stdout, line)
}

func read(r io.Reader) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		log.Print(err)
	}
	line = string(data)
}

func print(w io.Writer, line string) {
	b := bufio.NewWriter(w)
	b.WriteString(line)
	b.Flush()
}
