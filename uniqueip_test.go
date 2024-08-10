package main

import (
	"testing"
)

const SAMPLE_LINE_COUNT = 10

func BenchmarkLineCount(b *testing.B) {
	getUniqueAddresses("ip_addresses_sample_mini.txt")
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
