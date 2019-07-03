package main

import (
	"fmt"
	"github.com/fatih/color"
	"jaytaylor.com/html2text"
	"net/http"
	"os"
	"strings"
)

// declaring supported failed states
var failedState = []string{"NOK", "down", "degraded service"}

// some fancy output colors
var red = color.New(color.FgRed, color.Bold)
var green = color.New(color.FgHiGreen, color.Bold)

func main() {
	// checking if exactly one argument is supplied
	if len(os.Args) != 2 {
		fmt.Println("Too little or too many parameters supplied to the application... Please provide exactly one parameter: URL. Exiting...")
		os.Exit(1)
	}

	// getting bytes from the destination
	resp, err := http.Get(os.Args[1])
	if err != nil || resp == nil || resp.StatusCode != 200 {
		fmt.Println("Couldn't establish the connection to a status page, please check the correctness of provided URL")
		os.Exit(1)
	}
	defer resp.Body.Close()

	// transforming bytes into html into text
	text, err := html2text.FromReader(resp.Body, html2text.Options{PrettyTables: true})
	if err != nil {
		panic(err)
	}

	// a resulting string becomes a slice, for iteration reasons
	stringsArray := strings.Split(strings.TrimSpace(text), "\n")
	// splitting to get name of the service and it's state
	for _, string := range deleteEmpty(stringsArray) {
		s := strings.Split(string, ":")
		if stringInSlice(strings.TrimSpace(s[1]), failedState) {
			// actual
			red.Printf("The service '%v' is in '%v' state\n", s[0], strings.TrimSpace(s[1]))
		} else {
			green.Printf("The service '%v' is in '%v' state\n", s[0], strings.TrimSpace(s[1]))
		}
	}
}

// helper methods
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
