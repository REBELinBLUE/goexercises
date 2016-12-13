package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

var urls = []string{
	"https://www.google.co.uk",
	"http://golang.org",
	"http://www.bbc.co.uk",
	"https://uk.yahoo.com/",
}

type result struct {
	URL     string
	Elapsed time.Duration
}

func req(ctx context.Context, url string, back chan result) {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	start := time.Now()
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	resp.Body.Close()
	elapsed := time.Since(start)

	back <- result{
		URL:     url,
		Elapsed: elapsed / time.Millisecond,
	}
}

func makeRequests() {
	ctx, cancel := context.WithCancel(context.Background())
	back := make(chan result)
	fmt.Println("Making requests...")
	for _, url := range urls {
		go req(ctx, url, back)
	}

	for i := 0; i < len(urls); i++ {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		case u := <-back:
			fmt.Printf("%s took %d ms!\n", u.URL, u.Elapsed)
		}
	}

	cancel()
}

func main() {
	makeRequests()
}
