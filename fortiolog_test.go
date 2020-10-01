package bench

import (
	"testing"

	log "fortio.org/fortio/log"
)

func BenchmarkFortiologTextNegative(b *testing.B) {
	stream := &blackholeStream{}
	log.SetLogLevel(log.Error)
	log.SetOutput(stream)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Infof("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(0) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkFortiologTextPositive(b *testing.B) {
	stream := &blackholeStream{}
	log.SetLogLevel(log.Info)
	log.SetOutput(stream)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Infof("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}
