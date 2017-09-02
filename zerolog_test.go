package bench

import (
	"testing"

	"github.com/rs/zerolog"
)

func BenchmarkZerologTextPositive(b *testing.B) {
	stream := &blackholeStream{}
	logger := zerolog.New(stream).With().Timestamp().Logger()
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

func BenchmarkZerologTextNegative(b *testing.B) {
	stream := &blackholeStream{}
	logger := zerolog.New(stream).
		Level(zerolog.ErrorLevel).
		With().Timestamp().Logger()
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

func BenchmarkZerologJSONNegative(b *testing.B) {
	stream := &blackholeStream{}
	logger := zerolog.New(stream).
		Level(zerolog.ErrorLevel).
		With().Timestamp().Logger()
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

func BenchmarkZerologJSONPositive(b *testing.B) {
	stream := &blackholeStream{}
	logger := zerolog.New(stream).With().Timestamp().Logger()
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
