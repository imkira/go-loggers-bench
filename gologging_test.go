package bench

import (
	"testing"

	log "github.com/op/go-logging"
)

func BenchmarkGologgingTextFilterParallel(b *testing.B) {
	stream := &blackholeStream{}
	logger := log.MustGetLogger("")
	subBackend := log.NewLogBackend(stream, "", 0)
	formatter := log.MustStringFormatter("%{time:2006-01-02T15:04:05Z07:00} %{level} %{message}")
	backend := log.NewBackendFormatter(subBackend, formatter)
	leveled := log.AddModuleLevel(backend)
	leveled.SetLevel(log.ERROR, "")
	logger.SetBackend(leveled)
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

func BenchmarkGologgingTextParallel(b *testing.B) {
	stream := &blackholeStream{}
	logger := log.MustGetLogger("")
	subBackend := log.NewLogBackend(stream, "", 0)
	formatter := log.MustStringFormatter("%{time:2006-01-02T15:04:05Z07:00} %{level} %{message}")
	backend := log.NewBackendFormatter(subBackend, formatter)
	leveled := log.AddModuleLevel(backend)
	logger.SetBackend(leveled)
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
