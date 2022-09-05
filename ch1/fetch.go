package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		hasPrefix := strings.HasPrefix(url, "http://")
		if !hasPrefix {
			url = "http://" + url
		}
		fmt.Printf("url %s\n", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		_, readErr := io.Copy(os.Stdout, resp.Body)
		fmt.Printf("status %s\n", resp.Status)
		resp.Body.Close()
		if readErr != nil {
			fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, readErr)
			os.Exit(1)
		}

	}
}
