package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/headfirstgo/general"
)

//Page defines API of a Page
type Page struct {
	URL  string
	Size int
}

func getPageSize(url string, channel chan Page) {
	fmt.Printf("Getting %s\n", url)
	response, err := http.Get(url)
	general.HandleErr(err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	general.HandleErr(err)
	channel <- Page{url, len(body)}
}

func exerciseOne() {
	urls := []string{
		"https://example.com",
		"https://facebook.com",
		"https://google.com",
	}
	pages := make(chan Page)
	for _, url := range urls {
		go getPageSize(url, pages)
		p := <-pages
		fmt.Printf("The size of %s is %d bytes\n", p.URL, p.Size)
	}
}

func exerciseTwo() {
	chanA := make(chan int)
	chanB := make(chan int)
	go func(channel chan int) {
		channel <- 1
		channel <- 3
	}(chanA)

	go func(channel chan int) {
		channel <- 2
		channel <- 4
	}(chanB)
	fmt.Println(<-chanA)
	fmt.Println(<-chanB)
	fmt.Println(<-chanA)
	fmt.Println(<-chanB)
}

func main() {
	fmt.Println("Printing Size Pages...")
	exerciseOne()
	exerciseTwo()
}
