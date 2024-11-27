package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	listUrl := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.wikipedia.org",
		"https://www.twitter.com",
		"https://www.instagram.com",
		"https://www.linkedin.com",
		"https://www.reddit.com",
		"https://www.netflix.com",
		"https://www.apple.com",
		"https://www.microsoft.com",
		"https://www.bbc.com",
		"https://www.cnn.com",
		"https://www.nytimes.com",
		"https://www.weather.com",
		"https://www.forbes.com",
		"https://www.theguardian.com",
		"https://www.healthline.com",
		"https://www.imdb.com",
	}

	workerNum := 4

	// Atomic first then Mutex
	startTime1 := time.Now()
	fmt.Println("\n----------\nAtomic start")
	errorCountAtomic := testNClientAtomic(&listUrl, workerNum)
	ellapsedAtomic := time.Since(startTime1)

	startTime2 := time.Now()
	fmt.Println("\nMutex start")
	errorCountMutex := testNClientMutex(&listUrl, workerNum)
	ellapsedMutex := time.Since(startTime2)

	fmt.Printf("\nAtomic elapsed time %s, error count %d", ellapsedAtomic, errorCountAtomic)
	fmt.Printf("\nMutex elapsed time %s, error count %d", ellapsedMutex, errorCountMutex)

	// Mutex first then Atomic
	startTime1 = time.Now()
	fmt.Println("\n----------\nAtomic start")
	errorCountAtomic = testNClientAtomic(&listUrl, workerNum)
	ellapsedAtomic = time.Since(startTime1)

	startTime2 = time.Now()
	fmt.Println("\nMutex start")
	errorCountMutex = testNClientMutex(&listUrl, workerNum)
	ellapsedMutex = time.Since(startTime2)

	fmt.Printf("\nAtomic elapsed time %s, error count %d", ellapsedAtomic, errorCountAtomic)
	fmt.Printf("\nMutex elapsed time %s, error count %d", ellapsedMutex, errorCountMutex)
}

type safeCounter struct {
	mu    sync.Mutex
	count int32
}

func (counter *safeCounter) increase() {
	counter.mu.Lock()
	counter.count++
	counter.mu.Unlock()
}

func testNClientMutex(listUrl *[]string, clientNum int) int32 {
	clients := make([]http.Client, clientNum)
	for i := 0; i < clientNum; i++ {
		clients[i] = http.Client{Timeout: time.Second * 2}
	}

	inputChan := make(chan string)
	wg := sync.WaitGroup{}
	counter := safeCounter{}

	for i := 0; i < clientNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range inputChan {
				err := checkURL(&clients[i], strconv.Itoa(i), url)
				if err != nil {
					counter.increase()
				}
			}
		}()
	}

	for _, url := range *listUrl {
		inputChan <- url
	}
	close(inputChan)

	wg.Wait()
	return counter.count
}

func testNClientAtomic(listUrl *[]string, clientNum int) int32 {
	clients := make([]http.Client, clientNum)
	for i := 0; i < clientNum; i++ {
		clients[i] = http.Client{Timeout: time.Second * 2}
	}

	inputChan := make(chan string)
	wg := sync.WaitGroup{}
	var counter int32

	for i := 0; i < clientNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range inputChan {
				err := checkURL(&clients[i], strconv.Itoa(i), url)
				if err != nil {
					atomic.AddInt32(&counter, 1)
				}
			}
		}()
	}

	for _, url := range *listUrl {
		inputChan <- url
	}
	close(inputChan)

	wg.Wait()
	return counter
}

func checkURL(client *http.Client, clientName string, url string) error {
	startTime := time.Now()

	res, err := client.Get(url)

	elapsed := time.Since(startTime)
	if err != nil {
		fmt.Printf("Client %s | %s | elapsed %s | ERROR %s\n", clientName, url, elapsed, err.Error())
		return err
	} else {
		fmt.Printf("Client %s | %s | elapsed %s | STATUS %s\n", clientName, url, elapsed, res.Status)
		return nil
	}
}

// func testOneClient(listUrl *[]string) {
// 	// fmt.Println("hello")
// 	client := http.Client{
// 		Timeout: time.Second * 1,
// 	}
// 	startTime := time.Now()
// 	for _, url := range *listUrl {
// 		checkURL(&client, "1", url)
// 	}

// 	elapsed := time.Since(startTime)

// 	fmt.Printf("\nTotal elapsed time %s\n", elapsed)
// }
