package main

import (
	"flag"
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

func printHelp() {
	fmt.Println(`Usage:

    app <command> [arguments]

The commands are:

    version     print App version
    help        show this message
    cat         print text files
    wget        download a file from the web`)
}

func main() {
	catCmd := flag.NewFlagSet("cat", flag.ExitOnError)
	wgetCmd := flag.NewFlagSet("wget", flag.ExitOnError)
	switch os.Args[1] {
	case "version":
		printVersion()
	case "cat":
		catCmd.Parse(os.Args[2:])
		cat(catCmd.Args())
	case "wget":
		wgetCmd.Parse(os.Args[2:])
		wget(wgetCmd.Args()[0])
	case "help":
		fallthrough
	default:
		printHelp()
	}
}
