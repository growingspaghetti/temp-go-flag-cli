package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func cat(files []string) {
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		bytes, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(bytes))
	}
}
