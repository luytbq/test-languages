package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"test/file/shared"
)

func main() {

	filename := "numbers.txt"
	dataLen := 1000000
	var minValue int64 = 10000000
	var maxValue int64 = 10000000000

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	elapsedTime := shared.ElapsedTime(func() {
		for i := 0; i < dataLen; i++ {
			line := strconv.Itoa(int(randomInt(minValue, maxValue))) + "\n"
			file.Write([]byte(line))
		}
	})

	file.Close()

	log.Println("create file took: ", elapsedTime)
}

func randomInt(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}
