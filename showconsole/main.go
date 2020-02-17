// Copyright (c) 2020 by Gilbert Ramirez <gram@alumni.rice.edu>

package main

import (
	"fmt"
	"os"

	"github.com/gilramir/consolesize"
)

func main() {

	rect, err := consolesize.GetConsoleWidthHeight()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%d x %d\n", rect.Width, rect.Height)
}
