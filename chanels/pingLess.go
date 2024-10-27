package chanels

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func PingLess() {
	path := flag.String("file", "file.txt", "path to urls file")
	flag.Parse()
	file, err := os.ReadFile(*path)
	if err != nil {
		panic(err.Error())
	}
	urlSlice := strings.Split(string(file), "\n")

	respCh := make(chan int)
	errCh := make(chan error)

	for _, v := range urlSlice {
		go ping(v, respCh, errCh)
	}

	for range urlSlice {
		select {
		case errRes := <-errCh:
			fmt.Println(errRes)
		case respRes := <-respCh:
			fmt.Println(respRes)
		}
	}

}

func ping(url string, respCh chan int, errCh chan error) {
	resp, err := http.Get(url)
	if err != nil {
		errCh <- err
		return
	}
	respCh <- resp.StatusCode
}
