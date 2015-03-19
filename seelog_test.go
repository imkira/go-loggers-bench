package bench

import (
	"testing"

	log "github.com/cihub/seelog"
)

func BenchmarkSeelogTextFilterParallel(b *testing.B) {
	stream := &blackholeStream{}
	logger, err := log.LoggerFromWriterWithMinLevelAndFormat(stream, log.ErrorLvl, "%Time %Level %Msg")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		defer logger.Flush()
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(0) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkSeelogTextParallel(b *testing.B) {
	stream := &blackholeStream{}
	logger, err := log.LoggerFromWriterWithMinLevelAndFormat(stream, log.TraceLvl, "%Time %Level %Msg")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		defer logger.Flush()
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}
