package main

import (
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/datamatrix"
	flag "github.com/spf13/pflag"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const version = "1.0"

const usage = `Usage: dmtx [options] [file]

Options:
  -h, --help            display this help
  -s, --size int        size in pixel (default 200)
  -o, --output string   output file
  -V, --version         display version
`

func main() {

	var input string
	var output string
	var size int
	var info bool

	flag.StringVarP(&output,
		"output",
		"o",
		"",
		"Line separated file output")
	flag.IntVarP(&size, "size", "s", 600, "")
	flag.BoolP("help", "h", false, "")
	flag.Usage = func() {
		//flag.PrintDefaults()
		fmt.Print(usage)
	}
	flag.BoolVarP(&info, "version", "V", false, "")
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 && !info {
		flag.Usage()
		os.Exit(1)
	}

	if len(args) > 1 {
		fmt.Println("Too many arguments")
		os.Exit(1)
	}

	if info {
		fmt.Println("dmtx by Philipp Speck [Version " + version + "]")
		fmt.Println("Copyright (C) 2022 Typomedia Foundation.")
		os.Exit(0)
	}

	input = args[0]
	if input != "" {
		content, err := os.ReadFile(input)
		if err != nil {
			log.Fatal("Error when opening file: ", err)
		}

		code, err := datamatrix.Encode(string(content))
		if err != nil {
			fmt.Println("Error creating barcode:", err)
			return
		}

		code, _ = barcode.Scale(code, size, size)

		if output == "" {
			output = basename(input) + ".png"
		}

		file, err := os.Create(output)
		if err != nil {
			fmt.Println("Error saving barcode:", err)
			return
		}
		defer file.Close()
		err = png.Encode(file, code)
		if err != nil {
			return
		}
		os.Exit(0)
	}
}

func basename(file string) string {
	return strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
}
