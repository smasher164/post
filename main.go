// Usage:
//
// 		post /PATH
//
// Post creates a set of h2,h3,h4 tags for a blogpost. It is used to create a table of contents.
//
package main

import (
	"bufio"
	"flag"
	"fmt"
	"html"
	"log"
	"os"
	"strings"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: post /PATH \n")
	os.Exit(2)
}

func main() {
	log.SetPrefix("svglatex: ")
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()
	path := strings.TrimSpace(flag.Arg(0))
	if path == "" {
		usage()
	}
	if path[0] != '/' {
		usage()
	}

	scanner := bufio.NewScanner(os.Stdin)
	scan := func() string {
		scanner.Scan()
		s := scanner.Text()
		s = strings.TrimSpace(s)
		s = html.EscapeString(s)
		return s
	}
	title := scan()
	desc := scan()
	date := scan()
	if title == "" || date == "" || scanner.Err() != nil {
		usage()
	}
	fmt.Printf("<h2><a href=%q>%s</a></h2>", path, title)
	if desc != "" {
		fmt.Printf("<h3>%s</h3>", desc)
	}
	fmt.Printf("<h4>%s</h4>", date)
}
