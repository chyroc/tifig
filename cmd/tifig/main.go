package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Chyroc/tifig"
)

var (
	help   bool
	input  string
	output string
)

func init() {
	flag.BoolVar(&help, "help", false, "show help")
	flag.StringVar(&input, "input", "", "input heic file path")
	flag.StringVar(&output, "output", "", "output heic file path")
	flag.Parse()
}

func main() {
	if help || input == "" {
		fmt.Println(`tifig-go

tifig -input <input-heic-file-path> [-output <output-heic-file-path>]
`)
		return
	}

	output, err := tifig.Convert("/Users/chenyunpeng/Downloads/IMG_2670.HEIC", "")
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(output)
}
