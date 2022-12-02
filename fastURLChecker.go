package main

import (
	"fmt"
	"net/http"
)

type requestResult struct{
	url string
	status string
}

func main(){
	c := make(chan requestResult)
	results := make(map[string]string)
	urls := []string{
		"https://github.com/",
		"https://reddit.com",
		"https://amazon.com",
		"https://airbnb.com",
		"https://indeed.com",
	}

	for _, url := range urls{
		go checkURL(url, c)
	}

	for i := 0; i < len(urls); i++{
		result := <- c
		results[result.url] = result.status
	}

	for url, status := range results{
		fmt.Println(url, status)
	}
}

func checkURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "SUCCESS"

	if(err != nil || resp.StatusCode >= 400){
		status = "FAILED"
	} 

	c <- requestResult{url, status}
}
