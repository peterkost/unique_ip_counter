package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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

  addressCount := make(map[string]int)
  for fileScanner.Scan() {
    addressCount[fileScanner.Text()] = addressCount[fileScanner.Text()] + 1
  }
  readFile.Close()

  uniqueCount := 0
  for _, count := range addressCount {
    if count == 1 {
      uniqueCount++
    }
  }
  return  uniqueCount
}

func getIpIndex(ip string) uint32 {
  ipParts := strings.Split(ip, ".")

  a, _ := strconv.Atoi(ipParts[0])
  b, _ := strconv.Atoi(ipParts[1])
  c, _ := strconv.Atoi(ipParts[2])
  d, _ := strconv.Atoi(ipParts[3])

  return uint32(a*256*256*256 + b*256*256 + c*256 + d)

}
