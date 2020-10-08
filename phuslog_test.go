package bench

import (
	"testing"

	"github.com/phuslu/log"
)

func BenchmarkPhuslogTextPositive(b *testing.B) {
	stream := &blackholeStream{}
	logger := log.Logger{
		Writer: log.IOWriter{stream},
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().Msg("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkPhuslogTextNegative(b *testing.B) {
	stream := &blackholeStream{}
	logger := log.Logger{
		Level:  log.ErrorLevel,
		Writer: log.IOWriter{stream},
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().Msg("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(0) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkPhuslogJSONNegative(b *testing.B) {
	stream := &blackholeStream{}
	logger := log.Logger{
		Level:  log.ErrorLevel,
		Writer: log.IOWriter{stream},
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().
				Str("rate", "15").
				Int("low", 16).
				Float32("high", 123.2).
				Msg("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(0) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkPhuslogJSONPositive(b *testing.B) {
	stream := &blackholeStream{}
	logger := log.Logger{
		Writer: log.IOWriter{stream},
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().
				Str("rate", "15").
				Int("low", 16).
				Float32("high", 123.2).
				Msg("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}

