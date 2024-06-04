package main

import (
	"flag"
	"fmt"
	"wifi/internal/app"
)

func main() {
	file_name := flag.String("f", "no", "file to write to")
	flag.Parse()

	if *file_name == "no" {
		fmt.Println("Error - No file to write to specified!")
		return
	}
	file := *file_name + ".csv"
	app.Interaction(file)
}
