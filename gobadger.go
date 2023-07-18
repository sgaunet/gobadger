package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/narqo/go-badge"
)

func main() {
	var outputFile, color, title, value string
	flag.StringVar(&outputFile, "o", "badge.svg", "output file name")
	flag.StringVar(&color, "c", "#5272B4", "color of badge")
	flag.StringVar(&title, "t", "", "title")
	flag.StringVar(&value, "v", "", "Value for the title")
	flag.Parse()

	if title == "" || value == "" {
		fmt.Fprintln(os.Stderr, "title and value are required")
		flag.Usage()
		os.Exit(1)
	}

	f, err := os.Create(outputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	defer f.Close()

	badge, err := badge.RenderBytes(title, value, badge.Color(color))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	f.Write(badge)
}
