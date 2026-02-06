package main

import (
	"fmt"
	"time"
	"net/http"
	"bytes"
	"encoding/json"
	
)

type res struct {
			url string
			statusCode int
			err error
		}


func sendAlert(message string) {
	webhookURL := "https://discord.com/api/webhooks/1469363388743155996/pB4J2FC6E3fyMwOGlZXdKFDlrX6cw38pD2kv54Wqy4jXvdI7UbYZ-bHkfNVjKa1IipN0"

	payload := map[string]string{"content": message}
	jsondata, _ := json.Marshal(payload)

	http.Post(webhookURL, "application/json", bytes.NewBuffer(jsondata))
}

func checkHealth(url string, ch chan<- res){
	client := http.Client{
		Timeout: 30 * time.Second,
	}

    resp, err := client.Get(url)
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
				Alert := fmt.Sprintf("Error checking %s: %v\n", result.url, result.err)
				fmt.Println(Alert)
				sendAlert(Alert)
			} else {
				fmt.Printf("URL: %s, Status Code: %d\n", result.url, result.statusCode)
			}
		}
	}
}
