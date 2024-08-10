package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
  printLineCount()
}

func printLineCount() {
  filePath := os.Args[1]
  readFile, err := os.Open(filePath)

  if err != nil {
    fmt.Println(err)
    return
  }
  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)

  lineCount := 0
  for fileScanner.Scan() {
    lineCount++
    if lineCount % 100000000 == 0 {
      fmt.Println(lineCount)
    }
  }

  readFile.Close()

  fmt.Println("total lines: ", lineCount)
}
