package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"test/file/shared"
)

func main() {
	elapsedTime := shared.ElapsedTime(useWorker)

	log.Println("useWorker took: ", elapsedTime)
}

func worker(c chan uint64, wg *sync.WaitGroup, count *uint64) {
	defer wg.Done()
	for number := range c {
		if shared.IsPrime(number) {
			log.Printf("%d is prime number", number)
			atomic.AddUint64(count, 1)
		}
	}
}

func useWorker() {
	file, err := os.Open("./numbers.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numWorker := 10
	var wg sync.WaitGroup
	c := make(chan uint64, 100)

	scanner := bufio.NewScanner(file)

	var countPrime uint64

	for i := 0; i < numWorker; i++ {
		wg.Add(1)
		go worker(c, &wg, &countPrime)
	}

	for scanner.Scan() {
		number, err := strconv.ParseUint(scanner.Text(), 10, 64)

		if err != nil {
			continue
		}

		c <- number
	}

	if scanner.Err() != nil {
		log.Fatal(err)
	}

	close(c)
	wg.Wait()
	log.Printf("total prime number: %d", countPrime)
}
