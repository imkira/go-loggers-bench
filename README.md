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

## Take 2 (September 2017)

Two years later, same test ran with Go 1.9 on Mac OSX 10.12.6 using 2.7GHz
Intel Core i7 Macbook Pro with 16 GB 2133 MHz LPDDR3 memory.

Overall, zerolog performed the best by a substantial margin with a constant
0 allocations for both text and JSON (output is always JSON) with positive and
negative tests.

## benchstat

### TextPositive

| test                    | op time        | op alloc sz  | op alloc count |
|-------------------------|----------------|--------------|----------------|
| GokitTextPositive-4     |   442ns ± 4%   |    256B ± 0% |      4.00 ± 0% |
| GologgingTextPositive-4 |   628ns ± 1%   |    920B ± 0% |      17.0 ± 0% |
| Log15TextPositive-4     |  3.60µs ± 3%   |  1.12kB ± 0% |      24.0 ± 0% |
| LogrusTextPositive-4    |   665ns ± 2%   |    320B ± 0% |      15.0 ± 0% |
| SeelogTextPositive-4    |  2.18µs ± 1%   |    440B ± 0% |      11.0 ± 0% |
| ZerologTextPositive-4   | **130ns ± 4%** | **0.00B**    |    **0.00**    |

### TextNegative

| test                    | op time         | op alloc sz | op alloc count |
|-------------------------|-----------------|-------------|----------------|
| GokitTextNegative-4     |   16.7ns ± 1%   |  32.0B ± 0% |      1.00 ± 0% |
| GologgingTextNegative-4 |   61.4ns ± 1%   |   144B ± 0% |      3.00 ± 0% |
| Log15TextNegative-4     |    142ns ± 3%   |   128B ± 0% |      2.00 ± 0% |
| LogrusTextNegative-4    | **0.98ns ± 4%** |**0.00B**    |    **0.00**    |
| SeelogTextNegative-4    |   22.5ns ± 2%   |  48.0B ± 0% |      2.00 ± 0% |
| ZerologTextNegative-4   |   4.34ns ± 0%   |**0.00B**    |    **0.00**    |

### JSONPositive

| test                    | op time        | op alloc sz | op alloc count |
|-------------------------|----------------|-------------|----------------|
| GokitJSONPositive-4     |  1.42µs ± 4%   | 1.55kB ± 0% |      24.0 ± 0% |
| Log15JSONPositive-4     |  6.56µs ± 1%   | 2.01kB ± 0% |      30.0 ± 0% |
| LogrusJSONPositive-4    |  1.81µs ± 3%   | 2.45kB ± 0% |      33.0 ± 0% |
| ZerologJSONPositive-4   | **195ns ± 3%** |**0.00B**    |    **0.00**    |

### JSONNegative

| test                    | op time         | op alloc sz | op alloc count |
|-------------------------|-----------------|-------------|----------------|
| GokitJSONNegative-4     |   27.3ns ± 2%   |   128B ± 0% |      1.00 ± 0% |
| Log15JSONNegative-4     |    189ns ± 2%   |   320B ± 0% |      3.00 ± 0% |
| LogrusJSONNegative-4    |    257ns ± 2%   |   752B ± 0% |      5.00 ± 0% |
| ZerologJSONNegative-4   | **6.39ns ± 2%** |**0.00B**    |    **0.00**    |

## Raw data

### TextPositive

| test                             | ops      | ns/op         | bytes/op    | allocs/op       |
|----------------------------------|----------|---------------|-------------|-----------------|
| BenchmarkGokitTextPositive-4     | 20000000 |   428 ns/op   |  256 B/op   |   4 allocs/op   |
| BenchmarkGologgingTextPositive-4 | 10000000 |   621 ns/op   |  920 B/op   |  15 allocs/op   |
| BenchmarkLog15TextPositive-4     |  2000000 |  3612 ns/op   | 1120 B/op   |  24 allocs/op   |
| BenchmarkLogrusTextPositive-4    | 10000000 |   657 ns/op   |  320 B/op   |  10 allocs/op   |
| BenchmarkSeelogTextPositive-4    |  3000000 |  2197 ns/op   |  440 B/op   |  11 allocs/op   |
| BenchmarkZerologTextPositive-4   | 50000000 | **125 ns/op** |  **0 B/op** | **0 allocs/op** |

### TextNegative

| test                             | ops         | ns/op          | bytes/op    | allocs/op       |
|----------------------------------|-------------|----------------|-------------|-----------------|
| BenchmarkGokitTextNegative-4     |   500000000 |   16.7 ns/op   |   32 B/op   |   1 allocs/op   |
| BenchmarkGologgingTextNegative-4 |   100000000 |   60.8 ns/op   |  144 B/op   |   2 allocs/op   |
| BenchmarkLog15TextNegative-4     |    50000000 |    146 ns/op   |  128 B/op   |   1 allocs/op   |
| BenchmarkLogrusTextNegative-4    | 10000000000 | **1.02 ns/op** |    0 B/op   |   0 allocs/op   |
| BenchmarkSeelogTextNegative-4    |   300000000 |   22.1 ns/op   |   48 B/op   |   2 allocs/op   |
| BenchmarkZerologTextNegative-4   |  2000000000 |   4.34 ns/op   |  **0 B/op** | **0 allocs/op** |

### JSONPositive

| test                           | ops      | ns/op         | bytes/op    | allocs/op       |
|--------------------------------|----------|---------------|-------------|-----------------|
| BenchmarkGokitJSONPositive-4   |  5000000 |  1398 ns/op   | 1552 B/op   |  24 allocs/op   |
| BenchmarkLog15JSONPositive-4   |  1000000 |  6599 ns/op   | 2008 B/op   |  30 allocs/op   |
| BenchmarkLogrusJSONPositive-4  |  5000000 |  1761 ns/op   | 2450 B/op   |  33 allocs/op   |
| BenchmarkZerologJSONPositive-4 | 30000000 | **195 ns/op** |  **0 B/op** | **0 allocs/op** |

### JSONNegative

| test                           | ops        | ns/op          | bytes/op   | allocs/op       |
|--------------------------------|------------|----------------|------------|-----------------|
| BenchmarkGokitJSONNegative-4   |  300000000 |   27.0 ns/op   | 128 B/op   |   1 allocs/op   |
| BenchmarkLog15JSONNegative-4   |   30000000 |    188 ns/op   | 320 B/op   |   3 allocs/op   |
| BenchmarkLogrusJSONNegative-4  |   30000000 |    255 ns/op   | 752 B/op   |   5 allocs/op   |
| BenchmarkZerologJSONNegative-4 | 1000000000 | **6.26 ns/op** | **0 B/op** | **0 allocs/op** |
