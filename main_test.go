package main

import "testing"

func BenchmarkDecoder(b *testing.B) {
	b.StartTimer()
	Decoder()
	b.StopTimer()
}

func BenchmarkMarshal(b *testing.B) {
	b.StartTimer()
	Marshal()
	b.StopTimer()
}
