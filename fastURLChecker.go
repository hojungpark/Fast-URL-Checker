package main

import (
	"fmt"
	"net/http"
)

func main(){
	urls := []string{
		"https://github.com/",
		"https://reddit.com",
		"https://amazon.com",
		"https://airbnb.com",
		"https://indeed.com",
	}

	for _, url := range urls{
		checkURL(url)
	}
}

func checkURL(url string){
	resp, err := http.Get(url)
	if(err != nil || resp.StatusCode >= 400){
		fmt.Println(url, "FAILED")
	} else {
		fmt.Println(url, "SUCCESS")
	}
}
