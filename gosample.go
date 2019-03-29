package gosample

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Gets() {
	urls := []string{
		"http://example.com",
		"http://example.net",
		"http://example.org",
	}
	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			fmt.Println(url, res.Status)
		}(url)
	}
	// wait Goroutine
	time.Sleep(time.Second)
	fmt.Print("succeeded")
}
