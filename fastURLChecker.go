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
	results := make(map[string]string)
	urls := []string{
		"https://github.com/",
		"https://reddit.com",
		"https://amazon.com",
		"https://airbnb.com",
		"https://indeed.com",
	}

	for _, url := range urls{
		result := checkURL(url)
		results[result.url] = result.status
	}

	for url, status := range results{
		fmt.Println(url, status)
	}
}

func checkURL(url string) requestResult{
	resp, err := http.Get(url)
	status := "SUCCESS"

	if(err != nil || resp.StatusCode >= 400){
		status = "FAILED"
	} 

	return requestResult{url, status}
}
