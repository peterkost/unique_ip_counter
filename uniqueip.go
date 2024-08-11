package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"sync"
	"time"
)

const BUFFER_SIZE = 2048 * 2048
const NUM_WORKERS = 10
const CHANNEL_BUFFER = 10

var seenIps = make([]bool, math.MaxUint32)

func main() {
  started := time.Now()
	filePathPtr := flag.String("f", "", "Path to ip address file")
	flag.Parse()
  res := getUniqueAddresses(*filePathPtr)
  fmt.Println("Unique IPs: ", res)
	fmt.Printf("%0.6f\n", time.Since(started).Seconds())
}

func consumer(input chan []byte, wg *sync.WaitGroup) {
	defer wg.Done()

	var ip [4]uint32
	for ipBytes := range input {
		ipIndex := 0

		for _, b := range ipBytes {
			if b >= 48 && b <= 57 {
				ip[ipIndex] *= 10
				ip[ipIndex] += uint32(b - '0')
			} else if b == '.' {
				ipIndex++
			} else if b == '\n' {
				seenIpIndex := getIpIndex(ip)
				seenIps[seenIpIndex] = true

				ipIndex = 0
				ip = [4]uint32{}
			}

		}
	}

	var emptyArray [4]uint32
	if ip != emptyArray {
		seenIpIndex := getIpIndex(ip)
		seenIps[seenIpIndex] = true
	}
}

func getUniqueAddresses(filePath string) int {
	seenIps = make([]bool, math.MaxUint32)

	inputChannels := make([]chan []byte, NUM_WORKERS)

	var wg sync.WaitGroup
	wg.Add(NUM_WORKERS)

	for i := 0; i < NUM_WORKERS; i++ {
		input := make(chan []byte, CHANNEL_BUFFER)

		go consumer(input, &wg)

		inputChannels[i] = input
	}

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	readBuffer := make([]byte, BUFFER_SIZE)
	leftoverBuffer := make([]byte, 1024)
	leftoverSize := 0
	currentWorker := 0
	for {
		numReadBytes, err := file.Read(readBuffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		lastNewLineIndex := 0
		for i := numReadBytes - 1; i >= 0; i-- {
			if readBuffer[i] == 10 {
				lastNewLineIndex = i
				break
			}
		}

		data := make([]byte, lastNewLineIndex+leftoverSize)
		copy(data, leftoverBuffer[:leftoverSize])
		copy(data[leftoverSize:], readBuffer[:lastNewLineIndex])
		copy(leftoverBuffer, readBuffer[lastNewLineIndex+1:numReadBytes])
		leftoverSize = numReadBytes - lastNewLineIndex - 1

		inputChannels[currentWorker] <- data

		currentWorker++
		if currentWorker >= NUM_WORKERS {
			currentWorker = 0
		}
	}

	for i := 0; i < NUM_WORKERS; i++ {
		close(inputChannels[i])
	}

	wg.Wait()

	uniqueIps := 0
	for _, seen := range seenIps {
		if seen {
			uniqueIps++
		}
	}
	return uniqueIps
}

func getIpIndex(ip [4]uint32) uint32 {
	return ip[0]*256*256*256 + ip[1]*256*256 + ip[2]*256 + ip[3]
}
