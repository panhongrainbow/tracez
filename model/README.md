# Performance Enhancement

I have to come up with strategies to improve `the performance of TracingData`.

The explanations are listed below:

## Omit some Json fields 

The TracingData data is very large, but I found that if `some fields can be Omited`, the performance can be improved.

Run benchmark tests to see the performance difference

```bash
$ cd tracez/model

$ go test -v -bench='^\QBenchmark_Estimate_omitemptySample' -run=none .
# goos: linux
# goarch: amd64
# pkg: github.com/panhongrainbow/tracez/model
# cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
# Benchmark_Estimate_omitemptySample
# Benchmark_Estimate_omitemptySample-8 821595 1406 ns/op # <<<<< faster
# PASS
# ok      github.com/panhongrainbow/tracez/model  1.829s

$ go test -v -bench='^\QBenchmark_Estimate_nonOmitemptySample' -run=none .
# goos: linux
# goarch: amd64
# pkg: github.com/panhongrainbow/tracez/model
# cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
# Benchmark_Estimate_nonOmitemptySample
# Benchmark_Estimate_nonOmitemptySample-8 636535 1820 ns/op # <<<<< slower
# PASS
# ok      github.com/panhongrainbow/tracez/model  1.973s
```

Anyway, when the time comes, I will figure out a way to omit the TracingData structure, but the original data needs to be backed up.