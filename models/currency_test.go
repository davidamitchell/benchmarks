package models

import (
	"testing"
)

func setup() {
	err := Setup()
	if err != nil {
		return
	}
}

func BenchmarkGetCurrencies(b *testing.B) {
	setup()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		GetCurrencies()
	}
}
