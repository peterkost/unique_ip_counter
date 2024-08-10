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
  getLineCount(*filePathPtr)
}

func getLineCount(filePath string) int {
  readFile, err := os.Open(filePath)

  if err != nil {
    fmt.Println(err)
    return 0
  }
  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)

  lineCount := 0
  for fileScanner.Scan() {
    lineCount++
  }

  readFile.Close()

  return lineCount
}
