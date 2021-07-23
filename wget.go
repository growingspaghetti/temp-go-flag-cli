package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func wget(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fileName := filepath.Base(url)
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if _, err := io.Copy(file, resp.Body); err != nil {
		log.Fatal(err)
	}
}
