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

I ran these tests on MacOSX 10.9.5 using a 4-Core 2Ghz Intel Core i7 MacBookPro
with 16GB GB 1600 MHz DDR3 memory.

Overall, logrus performed the best.

I was surprised with log15 because it seemed
not to take advantage of the 2 or 4 goroutines. I believe this is because it
probably performs most of its operations inside some mutext-protected execution
path right before flushing the log to the output stream.

When it comes to negative tests for JSON, both logrus and log15 made too many
allocations for my taste. This is especially bad if you carpet bomb your
program with debug-level log lines but want to disable it in production: it
puts too much unecessary pressure on GC, leads to all sorts of problems
including memory fragmentation. This is something I would like to see improved
in future versions!

## logrus (Text)

|test|ops|ns/op|bytes/Op|allocs/op|
|----|---|-----|--------|---------|
|BenchmarkLogrusTextPositive|1000000|5738|915|18|
|BenchmarkLogrusTextPositive-2|2000000|3718|920|18|
|BenchmarkLogrusTextPositive-4|3000000|2187|924|18|
|BenchmarkLogrusTextNegative|100000000|71.7|16|1|
|BenchmarkLogrusTextNegative-2|200000000|48.3|16|1|
|BenchmarkLogrusTextNegative-4|300000000|29.9|16|1|

## log15 (Text)

|test|ops|ns/op|bytes/Op|allocs/op|
|----|---|-----|--------|---------|
|BenchmarkLog15TextPositive|1000000|7437|1125|24|
|BenchmarkLog15TextPositive-2|1000000|8782|1129|24|
|BenchmarkLog15TextPositive-4|1000000|8816|1136|24|
|BenchmarkLog15TextNegative|10000000|909|128|1|
|BenchmarkLog15TextNegative-2|20000000|521|128|1|
|BenchmarkLog15TextNegative-4|20000000|316|128|1|

## go-logging (Text)

|test|ops|ns/op|bytes/Op|allocs/op|
|----|---|-----|--------|---------|
|BenchmarkGologgingTextPositive|2000000|4010|842|14|
|BenchmarkGologgingTextPositive-2|3000000|2766|848|14|
|BenchmarkGologgingTextPositive-4|5000000|1805|853|14|
|BenchmarkGologgingTextNegative|20000000|394|144|1|
|BenchmarkGologgingTextNegative-2|30000000|253|144|1|
|BenchmarkGologgingTextNegative-4|50000000|178|144|1|

## seelog (Text)

|test|ops|ns/op|bytes/Op|allocs/op|
|----|---|-----|--------|---------|
|BenchmarkSeelogTextPositive|2000000|3855|442|12|
|BenchmarkSeelogTextPositive-2|1000000|5137|444|12|
|BenchmarkSeelogTextPositive-4|2000000|4630|447|12|
|BenchmarkSeelogTextNegative|20000000|376|64|3|
|BenchmarkSeelogTextNegative-2|20000000|313|64|3|
|BenchmarkSeelogTextNegative-4|20000000|377|64|3|

## logrus (JSON)

|test|ops|ns/op|bytes/Op|allocs/op|
|----|---|-----|--------|---------|
|BenchmarkLogrusJSONPositive|500000|13055|2532|49|
|BenchmarkLogrusJSONPositive-2|1000000|8256|2543|49|
|BenchmarkLogrusJSONPositive-4|2000000|4831|2558|49|
|BenchmarkLogrusJSONNegative|3000000|2677|896|11|
|BenchmarkLogrusJSONNegative-2|5000000|1704|896|11|
|BenchmarkLogrusJSONNegative-4|10000000|1164|896|11|

## log15 (JSON)

|test|ops|ns/op|bytes/Op|allocs/op|
|----|---|-----|--------|---------|
|BenchmarkLog15JSONPositive|500000|12213|2025|47|
|BenchmarkLog15JSONPositive-2|500000|13813|2027|47|
|BenchmarkLog15JSONPositive-4|500000|14067|2029|47|
|BenchmarkLog15JSONNegative|5000000|1732|392|9|
|BenchmarkLog15JSONNegative-2|10000000|1095|392|9|
|BenchmarkLog15JSONNegative-4|10000000|659|392|9|
