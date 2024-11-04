package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// links := []string{
	// 	"http://google.com",
	// 	"http://facebook.com",
	// 	"http://stackoverflow.com",
	// 	"http://golang.org",
	// 	"http://amazon.com",
	// }
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
}
