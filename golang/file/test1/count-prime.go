package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"test/file/shared"
)

func main() {
	elapsedTime := shared.ElapsedTime(lineByLine)

	log.Println("lineByLine took: ", elapsedTime)
}

func lineByLine() {
	file, err := os.Open("./numbers.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var countPrime uint64
	for scanner.Scan() {
		number, err := strconv.ParseUint(scanner.Text(), 10, 64)

		if err != nil {
			continue
		}

		if shared.IsPrime(number) {
			log.Printf("%d is prime number", number)
			countPrime++
		}
	}

	if scanner.Err() != nil {
		log.Fatal(err)
	}

	log.Printf("total prime number: %d", countPrime)
}
