package bench

import (
	"testing"

	log "github.com/Sirupsen/logrus"
)

func BenchmarkLogrusTextParallel(b *testing.B) {
	stream := &blackholeStream{}
	logger := log.New()
	logger.Formatter = &log.TextFormatter{
		DisableColors:  true,
		FullTimestamp:  true,
		DisableSorting: true,
	}
	logger.Out = stream
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

func BenchmarkLogrusTextFilterParallel(b *testing.B) {
	stream := &blackholeStream{}
	logger := log.New()
	logger.Level = log.ErrorLevel
	logger.Formatter = &log.TextFormatter{
		DisableColors:  true,
		FullTimestamp:  true,
		DisableSorting: true,
	}
	logger.Out = stream
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

func BenchmarkLogrusJSONFilterParallel(b *testing.B) {
	stream := &blackholeStream{}
	logger := log.New()
	logger.Level = log.ErrorLevel
	logger.Formatter = &log.JSONFormatter{}
	logger.Out = stream
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.WithFields(log.Fields{
				"rate": "15",
				"low":  16,
				"high": 123.2,
			}).Info("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(0) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkLogrusJSONParallel(b *testing.B) {
	stream := &blackholeStream{}
	logger := log.New()
	logger.Formatter = &log.JSONFormatter{}
	logger.Out = stream
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.WithFields(log.Fields{
				"rate": "15",
				"low":  16,
				"high": 123.2,
			}).Info("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}
