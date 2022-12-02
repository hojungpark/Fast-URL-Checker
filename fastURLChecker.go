package main

import (
	"fmt"
	"net/http"
	"time"
)

type requestResult struct{
	url string
	status string
}

func main(){
	results := make(map[string]string)
	urls := []string{
		"https://github.com/",
		"https://reddit.com",
		"https://amazon.com",
		"https://airbnb.com",
		"https://indeed.com",
	}

	for _, url := range urls{
		go checkURL(url)
	}

	for url, status := range results{
		fmt.Println(url, status)
	}

	time.Sleep(time.Second*10)
}

func checkURL(url string) {
	resp, err := http.Get(url)
	status := "SUCCESS"

	if(err != nil || resp.StatusCode >= 400){
		status = "FAILED"
	} 

	fmt.Println(requestResult{url, status})
}
