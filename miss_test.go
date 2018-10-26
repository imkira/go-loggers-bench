package bench

import (
	"testing"

	"github.com/xgfone/miss"
)

func BenchmarkMissTextNegative(b *testing.B) {
	stream := &blackholeStream{}
	conf := miss.EncoderConfig{IsTime: true, IsLevel: true}
	encoder := miss.KvTextEncoder(stream, conf)
	logger := miss.New(encoder).Level(miss.ERROR)

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

func BenchmarkMissTextPositive(b *testing.B) {
	stream := &blackholeStream{}
	conf := miss.EncoderConfig{IsTime: true, IsLevel: true}
	encoder := miss.KvTextEncoder(stream, conf)
	logger := miss.New(encoder)

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

func BenchmarkMissJSONNegative(b *testing.B) {
	stream := &blackholeStream{}
	conf := miss.EncoderConfig{IsTime: true, IsLevel: true}
	encoder := miss.KvSimpleJSONEncoder(stream, conf)
	logger := miss.New(encoder).Level(miss.ERROR)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog", "rate", 15, "low", 16, "high", 123.2)
		}
	})

	if stream.WriteCount() != uint64(0) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkMissJSONPositive(b *testing.B) {
	stream := &blackholeStream{}
	conf := miss.EncoderConfig{IsTime: true, IsLevel: true}
	encoder := miss.KvSimpleJSONEncoder(stream, conf)
	logger := miss.New(encoder)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog", "rate", 15, "low", 16, "high", 123.2)
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}
