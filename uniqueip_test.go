package main

import (
	"testing"
)

func BenchmarkLineCount(b *testing.B) {
	getLineCount("ip_addresses_sample.txt")
}
