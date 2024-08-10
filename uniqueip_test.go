package main

import (
	"fmt"
	"math"
	"testing"
)

const SAMPLE_LINE_COUNT = 10

func BenchmarkLineCount(b *testing.B) {
  fmt.Println("start bench")
	getUniqueAddresses("input/ip_addresses_sample.txt")
}

func TestThreeUnique(t *testing.T) {
    result := getUniqueAddresses("input/three_unique.txt")
    expected := 3

    if result != expected {
        t.Errorf("Expected %d but got %d", expected, result)
    }
}

func TestAllUnique(t *testing.T) {
    result := getUniqueAddresses("input/all_unique.txt")
    expected := SAMPLE_LINE_COUNT

    if result != expected {
        t.Errorf("Expected %d but got %d", expected, result)
    }
}

func TestOneDupe(t *testing.T) {
    result := getUniqueAddresses("input/one_dupe.txt")
    expected := SAMPLE_LINE_COUNT - 1

    if result != expected {
        t.Errorf("Expected %d but got %d", expected, result)
    }
}

func TestAllSame(t *testing.T) {
    result := getUniqueAddresses("input/all_same.txt")
    expected :=  0

    if result != expected {
        t.Errorf("Expected %d but got %d", expected, result)
    }
}

func TestFirstIpIndex(t *testing.T) {
    result := getIpIndex("0.0.0.0")
    expected :=  uint32(0)

    if result != expected {
        t.Errorf("Expected %d but got %d", expected, result)
    }
}

func TestSecondIpIndex(t *testing.T) {
    result := getIpIndex("0.0.0.1")
    expected :=  uint32(1)

    if result != expected {
        t.Errorf("Expected %d but got %d", expected, result)
    }
}

func TestLastIpIndex(t *testing.T) {
    result := getIpIndex("255.255.255.255")
    expected := uint32(math.MaxUint32)

    if result != expected {
        t.Errorf("Expected %d but got %d", expected, result)
    }
}
