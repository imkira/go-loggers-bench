// SPDX-License-Identifier: MIT

package bench

import (
	"testing"

	"github.com/issue9/logs/v4"
)

func BenchmarkLogsTextPositive(b *testing.B) {
	stream := &blackholeStream{}
	logger := logs.New(logs.NewTextWriter("01-02-15:04:05", stream))

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkLogsTextNegative(b *testing.B) {
	stream := &blackholeStream{}
	logger := logs.New(logs.NewTextWriter("01-02-15:04:05", stream))
	logger.Enable(logs.LevelError)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(0) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkLogsJSONNegative(b *testing.B) {
	stream := &blackholeStream{}
	logger := logs.New(logs.NewJSONWriter(logs.MicroLayout, stream))
	logger.Enable(logs.LevelError)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.INFO().With("rate", "15").With("low", 16).With("high", 123.2).Print("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(0) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkLogsJSONPositive(b *testing.B) {
	stream := &blackholeStream{}
	logger := logs.New(logs.NewJSONWriter(logs.MicroLayout, stream))

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.INFO().With("rate", "15").With("low", 16).With("high", 123.2).Print("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}
