package main

import (
	"math"
	"testing"
)

func BenchmarkLineCount(b *testing.B) {
	getUniqueIpCount("input/ip_addresses_sample.txt")
}

func TestThreeDupes(t *testing.T) {
	result := getUniqueIpCount("input/eight.txt")
	expected := 8

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestAllUnique(t *testing.T) {
	result := getUniqueIpCount("input/ten.txt")
	expected := 10

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestOneDupe(t *testing.T) {
	result := getUniqueIpCount("input/ten_w_dupe.txt")
	expected := 10

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestAllSame(t *testing.T) {
	result := getUniqueIpCount("input/one.txt")
	expected := 1

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestFirstIpIndex(t *testing.T) {
	result := getIpIndex([4]uint32{0, 0, 0, 0})
	expected := uint32(0)

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestSecondIpIndex(t *testing.T) {
	result := getIpIndex([4]uint32{0, 0, 0, 1})
	expected := uint32(1)

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestLastIpIndex(t *testing.T) {
	result := getIpIndex([4]uint32{255, 255, 255, 255})
	expected := uint32(math.MaxUint32)

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
