package main

import (
	"log"
	"net/http"
	"sync"
)

func main() {
	code := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			getHttpCode(code)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(code)
	}()

	for res := range code {
		log.Println(res)
	}
}

func getHttpCode(codeCh chan int) {
	req, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		log.Println(err.Error())
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := new(http.Client).Do(req)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()

	codeCh <- resp.StatusCode
}

/** Kol c wg.Wait() */

// func main() {
// 	t := time.Now()
// 	var wg sync.WaitGroup
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go func() {
// 			reqHadler(i)
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	log.Println(time.Since(t))
// }

// func reqHadler(num int) {
// 	req, err := http.NewRequest("GET", "https://google.com", nil)
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	req.Header.Set("User-Agent", "Mozilla/5.0")

// 	resp, err := new(http.Client).Do(req)
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	defer resp.Body.Close()

// 	log.Println(resp.StatusCode, "  - ", strconv.Itoa(num))
// }
