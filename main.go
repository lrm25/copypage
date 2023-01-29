package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var page string

func init() {
	flag.StringVar(&page, "page", "", "HTML page to copy")
}

func main() {
	flag.Parse()

	if page == "" {
		fmt.Println("HTML page location must be defined")
		return
	}

	response, err := http.Get(page)
	if err != nil {
		fmt.Println("Error retrieving HTML page: " + err.Error())
		return
	}
	defer response.Body.Close()

	htmlBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading HTML response body: " + err.Error())
		return
	}

	localFile, err := os.Create("test.html")
	if err != nil {
		fmt.Println("Error creating local file: " + err.Error())
		return
	}

	if _, err := localFile.Write(htmlBytes); err != nil {
		fmt.Println("Error writing HTML data to local file: " + err.Error())
		return
	}
}
