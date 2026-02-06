package main

import (
	"fmt"
	"net/http"
)

type res struct {
			url string
			statusCode int
			err error
		}


func checkHealth(url string, ch chan<- res){
    resp, err := http.Get(url)
    if err != nil {
        ch <- res{url: url, err: err}
        return 
    }
    defer resp.Body.Close()

    ch <- res{url: url, statusCode: resp.StatusCode, err: nil}
}


func main(){
	var n int

	for{
		fmt.Print("Enter the No. of url whose health you want to check (or '0' to quit):")
		fmt.Scanln(&n)
		if n==0 {
			break;
		}
		ex := make([]string, n)

		fmt.Println("Enter the URLs to check:")
		for i :=0;i<n;i++ {
			fmt.Printf("Enter URL %d: ", i+1)
            fmt.Scanln(&ex[i])
		}
		ch := make(chan res)

		for _, url := range ex {
			go checkHealth(url, ch)
		}

		for i:=0; i<len(ex); i++ {
			result := <- ch
			if result.err != nil {
				fmt.Printf("Error checking %s: %v\n", result.url, result.err)
			} else {
				fmt.Printf("URL: %s, Status Code: %d\n", result.url, result.statusCode)
			}
		}
	}
}
