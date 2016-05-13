# Benchmarking logging libraries for Go

I compared a varity of logging libraries for Go, and performed 2 types of tests
for 2 logging formats.

Types of test:

- Positive tests: everything will be logged because the baseline log level is
  set to lower or equal the level I am logging with.
- Negative tests: nothing will be logged because I set the baseline log level
  to ```ERROR``` but log wih ```INFO```. This is important, because the faster
  it bails-out, and the less pressure the library puts on the GC, the better!

Formats:

- Text: there are a few minor differences among loggers, but I tried to follow
  a ```Time Level Message``` format).
- JSON: few loggers support this but it is still interesting.

# Running the tests

On a terminal, just execute:

```shell
make
```

# Results

I ran these tests on Mac OSX 10.11.4 using 1.6 GHz Intel Core i5 Macbook Air
with 8 GB 1600 MHz DDR3 memory.

Overall, Go kit performed the best, though it has a quirky way of specifying
levels.

I was surprised with log15 because it seemed not to take advantage of the 2 or 4
goroutines. I believe this is because it probably performs most of its
operations inside some mutex-protected execution path right before flushing the
log to the output stream.

When it comes to negative tests for JSON, all loggers made too many allocations
for my taste. This is especially bad if you carpet bomb your program with
debug-level log lines but want to disable it in production: it puts too much
unecessary pressure on GC, leads to all sorts of problems including memory
fragmentation. This is something I would like to see improved in future
versions!

## benchstat

### TextPositive

| test                    | op time     | op alloc sz | op alloc count |
|-------------------------|-------------|-------------|----------------|
| GokitTextPositive-4     | 1.65µs ± 4% |   800B ± 0% |      10.0 ± 0% |
| GologgingTextPositive-4 | 2.11µs ± 6% |   952B ± 0% |      17.0 ± 0% |
| Log15TextPositive-4     | 8.05µs ±11% | 1.12kB ± 0% |      24.0 ± 0% |
| LogrusTextPositive-4    | 2.79µs ± 9% |   880B ± 0% |      17.0 ± 0% |
| SeelogTextPositive-4    | 4.92µs ± 8% |   456B ± 0% |      12.0 ± 0% |

### TextNegative

| test                    | op time     | op alloc sz | op alloc count |
|-------------------------|-------------|-------------|----------------|
| GokitTextNegative-4     |  123ns ± 7% |  64.0B ± 0% |      3.00 ± 0% |
| GologgingTextNegative-4 |  215ns ± 8% |   160B ± 0% |      3.00 ± 0% |
| Log15TextNegative-4     |  542ns ±26% |   128B ± 0% |      1.00 ± 0% |
| LogrusTextNegative-4    | 42.2ns ± 2% |  16.0B ± 0% |      1.00 ± 0% |
| SeelogTextNegative-4    |  132ns ±13% |  64.0B ± 0% |      3.00 ± 0% |

### JSONPositive

| test                    | op time     | op alloc sz | op alloc count |
|-------------------------|-------------|-------------|----------------|
| GokitJSONPositive-4     | 5.13µs ±16% | 1.33kB ± 0% |      36.0 ± 0% |
| Log15JSONPositive-4     | 12.6µs ±20% | 1.80kB ± 0% |      40.0 ± 0% |
| LogrusJSONPositive-4    | 6.15µs ±20% | 2.35kB ± 0% |      43.0 ± 0% |

### JSONNegative

| test                    | op time     | op alloc sz | op alloc count |
|-------------------------|-------------|-------------|----------------|
| GokitJSONNegative-4     |  377ns ±16% |   232B ± 0% |      9.00 ± 0% |
| Log15JSONNegative-4     |  909ns ± 5% |   392B ± 0% |      9.00 ± 0% |
| LogrusJSONNegative-4    | 1.14µs ±24% |   832B ± 0% |      10.0 ± 0% |

## Raw data

### TextPositive

| test                             | ops     | ns/op      | bytes/op  | allocs/op    |
|----------------------------------|---------|------------|-----------|--------------|
| BenchmarkGokitTextPositive-4     | 1000000 | 1406 ns/op |  352 B/op |  7 allocs/op |
| BenchmarkGologgingTextPositive-4 | 1000000 | 2021 ns/op |  952 B/op | 17 allocs/op |
| BenchmarkLog15TextPositive-4     |  200000 | 7688 ns/op | 1120 B/op | 24 allocs/op |
| BenchmarkLogrusTextPositive-4    | 1000000 | 2679 ns/op |  880 B/op | 17 allocs/op |
| BenchmarkSeelogTextPositive-4    |  300000 | 4326 ns/op |  456 B/op | 12 allocs/op |

### TextNegative

| test                             | ops      | ns/op       | bytes/op | allocs/op   |
|----------------------------------|----------|-------------|----------|-------------|
| BenchmarkGokitTextNegative-4     | 20000000 | 119 ns/op   |  64 B/op | 3 allocs/op |
| BenchmarkGologgingTextNegative-4 | 10000000 | 210 ns/op   | 160 B/op | 3 allocs/op |
| BenchmarkLog15TextNegative-4     |  3000000 | 510 ns/op   | 128 B/op | 1 allocs/op |
| BenchmarkLogrusTextNegative-4    | 50000000 |  45.3 ns/op |  16 B/op | 1 allocs/op |
| BenchmarkSeelogTextNegative-4    | 10000000 | 116 ns/op   |  64 B/op | 3 allocs/op |

### JSONPositive

| test                          | ops    | ns/op       | bytes/op  | allocs/op    |
|-------------------------------|--------|-------------|-----------|--------------|
| BenchmarkGokitJSONPositive-4  | 300000 |  4804 ns/op | 1328 B/op | 36 allocs/op |
| BenchmarkLog15JSONPositive-4  | 200000 | 11835 ns/op | 1800 B/op | 40 allocs/op |
| BenchmarkLogrusJSONPositive-4 | 300000 |  5853 ns/op | 2345 B/op | 43 allocs/op |

### JSONNegative

| test                          | ops     | ns/op      | bytes/op | allocs/op    |
|-------------------------------|---------|------------|----------|--------------|
| BenchmarkGokitJSONNegative-4  | 5000000 |  328 ns/op | 232 B/op |  9 allocs/op |
| BenchmarkLog15JSONNegative-4  | 2000000 |  895 ns/op | 392 B/op |  9 allocs/op |
| BenchmarkLogrusJSONNegative-4 | 1000000 | 1015 ns/op | 832 B/op | 10 allocs/op |
