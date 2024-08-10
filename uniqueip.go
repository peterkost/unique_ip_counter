package main

import (
	"bufio"
	"fmt"
	"os"
  "flag"
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
