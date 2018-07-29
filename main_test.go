package main

import "testing"

func BenchmarkDecoder(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Decoder()
	}
	b.StopTimer()
}

func BenchmarkMarshal(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Marshal()
	}
	b.StopTimer()
}
