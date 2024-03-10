package main

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/otiai10/gosseract/v2"

	"flag"
)

func main() {

	myFigure := figure.NewColorFigure("OCR CLI", "", "cyan", true)
	myFigure.Print()
	fmt.Println(color.YellowString(" By Federico Juretich <fedejuret@gmail.com>"))
	fmt.Println()

	image := flag.String("image", "", "Path to image")
	lang := flag.String("lang", "eng", "Image language")
	flag.Parse()

	if len(*image) == 0 {
		fmt.Println("Please complete with -image flag")
		os.Exit(1)
	}

	client := gosseract.NewClient()
	defer client.Close()

	client.SetLanguage(*lang)
	client.SetImage(*image)
	text, err := client.Text()

	if err != nil {
		panic(err)
	}

	fmt.Println(color.GreenString("- Result -"))
	fmt.Println()

	fmt.Println(text)
}
