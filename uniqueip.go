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
const NUM_WORKERS = 25
const CHANNEL_BUFFERS = 25

var seenIps []bool

func main() {
	started := time.Now()

	filePathPtr := flag.String("f", "", "Path to ip address file")
	flag.Parse()

	uniqueIps := getUniqueIpCount(*filePathPtr)

	fmt.Println("Unique IPs: ", uniqueIps)
	fmt.Printf("%0.6f\n", time.Since(started).Seconds())
}

func getUniqueIpCount(filePath string) int {
	seenIps = make([]bool, math.MaxUint32 + 1)

	inputChannels := make([]chan []byte, NUM_WORKERS)

	var wg sync.WaitGroup
	wg.Add(NUM_WORKERS)

	for i := 0; i < NUM_WORKERS; i++ {
		input := make(chan []byte, CHANNEL_BUFFERS)

		go markIpsSeenInChunk(input, &wg)

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

func markIpsSeenInChunk(chunk chan []byte, wg *sync.WaitGroup) {
	defer wg.Done()

	var ipOctets [4]uint32
	var octetIndex int

	for bytes := range chunk {
		octetIndex = 0
		for _, b := range bytes {
			if b >= 48 && b <= 57 {
				ipOctets[octetIndex] *= 10
				ipOctets[octetIndex] += uint32(b - '0')
			} else if b == '.' {
				octetIndex++
			} else if b == '\n' {
				seenIpIndex := getIpIndex(ipOctets)
				seenIps[seenIpIndex] = true

				octetIndex = 0
				ipOctets = [4]uint32{}
			}
		}
	}

	// edge case where file doesn't end with new line
	if octetIndex != 0 {
		seenIpIndex := getIpIndex(ipOctets)
		seenIps[seenIpIndex] = true
	}
}

func getIpIndex(ip [4]uint32) uint32 {
	return ip[0]*256*256*256 + ip[1]*256*256 + ip[2]*256 + ip[3]
}
