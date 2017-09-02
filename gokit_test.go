package bench

import (
	"sync"
	"testing"

	"github.com/go-kit/kit/log"
)

// Go kit's logger has no concept of dynamically mutable levels. The idiom is to
// predeclare your desired level during construction. This is an example helper
// constructor that performs that work. If positive is true, both info and error
// are logged. Otherwise, only error is logged.
func newLeveledLogger(logger log.Logger, positive bool) *leveledLogger {
	infoLogger := log.NewNopLogger()
	if positive {
		infoLogger = log.With(logger, "level", "info")
	}
	return &leveledLogger{
		Info:  infoLogger,
		Error: log.With(logger, "level", "error"),
	}
}

type leveledLogger struct {
	Info  log.Logger
	Error log.Logger
}

// For now, manually synchronize writes to the stream.
type synchronizedStream struct {
	mtx sync.Mutex
	blackholeStream
}

func (s *synchronizedStream) Write(p []byte) (int, error) {
	s.mtx.Lock()
	n, err := s.blackholeStream.Write(p)
	s.mtx.Unlock()
	return n, err
}

func BenchmarkGokitJSONPositive(b *testing.B) {
	stream := &synchronizedStream{}
	logger := log.With(log.NewJSONLogger(stream), "ts", log.DefaultTimestampUTC)
	lvllog := newLeveledLogger(logger, true)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lvllog.Info.Log("msg", "The quick brown fox jumps over the lazy dog", "rate", 15, "low", 16, "high", 123.2)
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkGokitJSONNegative(b *testing.B) {
	stream := &synchronizedStream{}
	logger := log.With(log.NewJSONLogger(stream), "ts", log.DefaultTimestampUTC)
	lvllog := newLeveledLogger(logger, false)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lvllog.Info.Log("msg", "The quick brown fox jumps over the lazy dog", "rate", 15, "low", 16, "high", 123.2)
		}
	})

	if stream.WriteCount() != uint64(0) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkGokitTextPositive(b *testing.B) {
	stream := &synchronizedStream{}
	logger := log.With(log.NewLogfmtLogger(stream), "ts", log.DefaultTimestampUTC)
	lvllog := newLeveledLogger(logger, true)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lvllog.Info.Log("msg", "The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkGokitTextNegative(b *testing.B) {
	stream := &synchronizedStream{}
	logger := log.With(log.NewLogfmtLogger(stream), "ts", log.DefaultTimestampUTC)
	lvllog := newLeveledLogger(logger, false)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lvllog.Info.Log("msg", "The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(0) {
		b.Fatalf("Log write count")
	}
}
