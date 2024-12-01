package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"http://www.amazon.com",
		"http://www.golang.org",
	}
	c := make(chan string)

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }
	var l = 1

	for l > 0 {
		for _, url := range links {
			go func(url string) {
				time.Sleep(5 * time.Second)
				checkLink(url, c)
			}(url)
		}
		time.Sleep(3 * time.Second)
	}
	// os.Exit(0)

}

func checkLink(url string, c chan string) {
	resp, err := http.Get(url)
	var msg string
	if err != nil {
		msg = fmt.Sprintf("http.Get method error: %v", err)
		fmt.Println(msg)
		c <- msg
		return
	}
	status := resp.StatusCode
	if status != 200 {
		msg = fmt.Sprintf("%s is not healthy. Status: %d", url, status)
	} else {
		msg = fmt.Sprintf("%s is healthy. Status: %d", url, status)
	}
	fmt.Println(msg)
	c <- msg
}
