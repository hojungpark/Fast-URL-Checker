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
	urls := []string{
		"https://github.com/",
		"https://reddit.com",
		"https://amazon.com",
		"https://airbnb.com",
		"https://indeed.com",
	}

	for _, url := range urls{
		result := checkURL(url)
		fmt.Println(result.url, result.status)
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
