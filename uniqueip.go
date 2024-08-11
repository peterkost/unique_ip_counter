package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	filePathPtr := flag.String("f", "", "Path to ip address file")
	flag.Parse()
	getUniqueAddresses(*filePathPtr)
}

func getUniqueAddresses(filePath string) int {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	br := bufio.NewReader(f)

	var previouslySeen [4294967296]bool
	total := 0
	duplicates := 0

	var ip [4]uint32
	ipIndex := 0

	for {
		b, err := br.ReadByte()

		if err != nil && !errors.Is(err, io.EOF) {
			fmt.Println(err)
			break
		}

		if b >= 48 && b <= 57 {
			ip[ipIndex] *= 10
			ip[ipIndex] += uint32(b - '0')
		} else if b == '.' {
			ipIndex++
		} else if b == '\n' {

			seenIpIndex := getIpIndex(ip)

			if previouslySeen[seenIpIndex] {
				duplicates++
			} else {
				previouslySeen[seenIpIndex] = true
			}

			total++

			ipIndex = 0
			ip = [4]uint32{}
		}

		if err != nil {
			break
		}
	}

	f.Close()

	return total - duplicates

}

func getIpIndex(ip [4]uint32) uint32 {
	return ip[0]*256*256*256 + ip[1]*256*256 + ip[2]*256 + ip[3]
}
