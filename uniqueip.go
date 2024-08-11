package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
)

func main() {
	filePathPtr := flag.String("f", "", "Path to ip address file")
	flag.Parse()
	getUniqueAddresses(*filePathPtr)
}

func getUniqueAddresses(filePath string) int {
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
		return 0
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var previouslySeen [4294967296]bool
	total := 0
	duplicates := 0
	for fileScanner.Scan() {
		ipAddress := fileScanner.Bytes()
		ipIndex := getIpIndex(ipAddress)

		if previouslySeen[ipIndex] {
			duplicates++
		}

		previouslySeen[ipIndex] = true
		total++
	}
	readFile.Close()

	return total - duplicates
}

func bytesToInt(ipChunk []byte) uint32 {
	var res uint32 = 0
	for _, ascii := range ipChunk {
		res *= 10
		res += uint32(ascii - '0')
	}
	return res
}

func getIpIndex(ip []byte) uint32 {
	firstPeriod := bytes.IndexByte(ip, byte('.'))
	firstNumber := bytesToInt(ip[:firstPeriod])

	secondPeriod := bytes.IndexByte(ip[firstPeriod+1:], byte('.')) + firstPeriod + 1
	secondNumber := bytesToInt(ip[firstPeriod+1 : secondPeriod])

	thirdPeriod := bytes.IndexByte(ip[secondPeriod+1:], byte('.')) + secondPeriod + 1
	thirdNumber := bytesToInt(ip[secondPeriod+1 : thirdPeriod])

	fourthNumber := bytesToInt(ip[thirdPeriod+1:])

	return firstNumber*256*256*256 + secondNumber*256*256 + thirdNumber*256 + fourthNumber

}

