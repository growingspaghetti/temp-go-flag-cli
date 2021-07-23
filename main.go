package main

import (
	"fmt"
	"os"
)

var (
	version  string
	commitId string
)

func printVersion() {
	fmt.Println("Version:", version)
	fmt.Println("Commit id:", commitId)
}

func main() {
	switch os.Args[1] {
	case "version":
		printVersion()
	}
}
