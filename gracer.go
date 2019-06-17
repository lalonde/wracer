package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	b = "https://www.bing.com/search?q="
	g = "https://www.google.com/search?q="
	y = "https://search.yahoo.com/search?p="
)

func main() {
	c := make(chan string, 1)
	r := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to gracer, the search engine racer\n")
	fmt.Print("Enter a search term: ")

	s, _ := r.ReadString('\n')
	s = url.QueryEscape(s)

	go query(y + s, c)
	go query(b + s, c)
	go query(g + s, c)

	fmt.Println("")
	fmt.Println("Winnner", <-c, "\n")
	fmt.Println("Second", <-c, "\n")
	fmt.Println("Third", <-c, "\n")
}

func query(u string, rc chan<- string) {
	s := time.Now()
	r, err := http.Get(u)
	if err != nil {
		rc <- fmt.Sprint("The race is flawed.\n\n", u, " errored:\n", err)
		return
	}
	defer r.Body.Close()
	t := time.Since(s).String()
	url, err := url.Parse(u)
	rc <- fmt.Sprint(url.Host, " in ", t)
}
