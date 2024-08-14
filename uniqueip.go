package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"time"
)

func main() {
	started := time.Now()

	filePathPtr := flag.String("f", "", "Path to ip address file")
	flag.Parse()

	uniqueIps := getUniqueIpCount(*filePathPtr)

	fmt.Println("Unique IPs: ", uniqueIps)
	fmt.Printf("%0.6f\n", time.Since(started).Seconds())
}

func getUniqueIpCount(filePath string) int {
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	br := bufio.NewReader(file)

	seenIps := make([]bool, math.MaxUint32+1)
	total := 0
	duplicates := 0

	var ip [4]uint32
	ipIndex := 0

	for {
		currentByte, err := br.ReadByte()

		if err != nil && !errors.Is(err, io.EOF) {
			fmt.Println(err)
			break
		}

		if currentByte >= 48 && currentByte <= 57 {
			ip[ipIndex] *= 10
			ip[ipIndex] += uint32(currentByte - '0')
		} else if currentByte == '.' {
			ipIndex++
		} else if currentByte == '\n' {

			seenIpIndex := getIpIndex(ip)

			if seenIps[seenIpIndex] {
				duplicates++
			} else {
				seenIps[seenIpIndex] = true
			}

			total++

			ipIndex = 0
			ip = [4]uint32{}
		}

		if err != nil {
			break
		}
	}

	return total - duplicates
}

func getIpIndex(ip [4]uint32) uint32 {
	return ip[0]*256*256*256 + ip[1]*256*256 + ip[2]*256 + ip[3]
}
