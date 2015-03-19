# Benchmarking logging libraries for Go

I performed 2 types of tests for 2 types of formats, against a variety of libraries.

Types:

- Positive match: everything will be logged.
- Negative match (tests containing the word ```Filter``` on its name): nothing
  will be logged because I set the baseline log level to ```ERROR``` but log
  wih ```INFO```. This is important, because the faster the library is able to
  "bail-out", the better.

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

Overall, logrus performs the best.

I was surprised with log15 because it seemed
not to take advantage of the 2 or 4 goroutines. I believe this is because it
probably performs most of its operations inside some mutext-protected execution
path right before flushing the log to the output stream.

When it comes to filtering, I thought logrus makes too many allocations when it
should just bail-out as fast as it can.

## logrus (Text)

|TestName|Ops|Nanoseconds/Op|Bytes/Op|Allocs/Op|
|--------|---|--------------|--------|---------|
|BenchmarkLogrusTextParallel|1000000|5779|915|18|
|BenchmarkLogrusTextParallel-2|2000000|3722|920|18|
|BenchmarkLogrusTextParallel-4|3000000|2245|924|18|
|BenchmarkLogrusTextFilterParallel|20000000|379|128|3|
|BenchmarkLogrusTextFilterParallel-2|30000000|246|128|3|
|BenchmarkLogrusTextFilterParallel-4|50000000|170|128|3|


## log15 (Text)

|TestName|Ops|Nanoseconds/Op|Bytes/Op|Allocs/Op|
|--------|---|--------------|--------|---------|
|BenchmarkLog15TextParallel|1000000|7674|1125|24|
|BenchmarkLog15TextParallel-2|1000000|8843|1129|24|
|BenchmarkLog15TextParallel-4|1000000|8868|1136|24|
|BenchmarkLog15TextFilterParallel|10000000|921|128|1|
|BenchmarkLog15TextFilterParallel-2|20000000|520|128|1|
|BenchmarkLog15TextFilterParallel-4|20000000|335|128|1|

## go-logging (Text)

|TestName|Ops|Nanoseconds/Op|Bytes/Op|Allocs/Op|
|--------|---|--------------|--------|---------|
|BenchmarkGologgingTextParallel|2000000|4050|842|14|
|BenchmarkGologgingTextParallel-2|3000000|2794|848|14|
|BenchmarkGologgingTextParallel-4|5000000|1836|853|14|
|BenchmarkGologgingTextFilterParallel|20000000|390|144|1|
|BenchmarkGologgingTextFilterParallel-2|30000000|253|144|1|
|BenchmarkGologgingTextFilterParallel-4|50000000|182|144|1|

## seelog (Text)

|TestName|Ops|Nanoseconds/Op|Bytes/Op|Allocs/Op|
|--------|---|--------------|--------|---------|
|BenchmarkSeelogTextParallel|2000000|3914|442|12|
|BenchmarkSeelogTextParallel-2|1000000|5253|444|12|
|BenchmarkSeelogTextParallel-4|2000000|4745|447|12|
|BenchmarkSeelogTextFilterParallel|20000000|379|64|3|
|BenchmarkSeelogTextFilterParallel-2|20000000|355|64|3|
|BenchmarkSeelogTextFilterParallel-4|20000000|379|64|3|

## logrus (JSON)

|TestName|Ops|Nanoseconds/Op|Bytes/Op|Allocs/Op|
|--------|---|--------------|--------|---------|
|BenchmarkLogrusJSONParallel|500000|12472|2532|49|
|BenchmarkLogrusJSONParallel-2|1000000|8437|2545|49|
|BenchmarkLogrusJSONParallel-4|1000000|5050|2559|49|
|BenchmarkLogrusJSONFilterParallel|3000000|2747|896|11|
|BenchmarkLogrusJSONFilterParallel-2|5000000|1722|896|11|
|BenchmarkLogrusJSONFilterParallel-4|10000000|1154|896|11|

## log15 (JSON)

|TestName|Ops|Nanoseconds/Op|Bytes/Op|Allocs/Op|
|--------|---|--------------|--------|---------|
|BenchmarkLog15JSONParallel|500000|12154|2025|47|
|BenchmarkLog15JSONParallel-2|500000|13911|2027|47
|BenchmarkLog15JSONParallel-4|500000|14130|2029|47|
|BenchmarkLog15JSONFilterParallel|5000000|1753|392|9|
|BenchmarkLog15JSONFilterParallel-2|10000000|1087|392|9|
|BenchmarkLog15JSONFilterParallel-4|10000000|665|392|9|
