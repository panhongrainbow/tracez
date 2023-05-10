# tracez

## My way to enforce OpenTelemetry

`Adding layers` to opentelemetry is like `the idea of breakpoints`.

When using `Intellij or Goland`, there is a `breakpoint feature`. Deeper and inner functions are skipped

Adding layers to opentelemetry has the effect of breakpoints, `but I made it into logs`.

Where should `the layer number be placed?` Each time entering a layer, the layer number increases by 1.

I `originally` planned to use `baggage` to write the layer number into ctx.

Why must baggage be used? To save trouble. If the `ctx withValue` function is used, there may be `race conditions`, and I will have to spend time carefully checking it.

Later I referred to [`rueidis's code`](https://github.com/redis/rueidis) and wanted to put the layer number in `the size position`.

`The layer number` becomes `part of the log`. See if `indexing` can be added to MongoDB, which is beneficial.

Let's see.

(我的目的是 让 span 有层数，让 span 可以合拼)

```go
func start(ctx context.Context, op string, size int, attrs []attribute.KeyValue) (context.Context, trace.Span) {
return tracer.Start(ctx, op, kind, attr(op, size), trace.WithAttributes(attrs...)) 
```

Provide more information later

(之后再补充)



## The application in throughput

(吞吐量的应要)

Some programs will `delay uploading data` until a certain number of records are accumulated and then upload them all at once.

The benefit of doing this is to `reduce CPU usage` and obtain better performance.

However, if the delay time is set `too high`, the `throughput will also decrease`.

So it is difficult to determine `how much the delay time should be set`.

Assume a program has `n goroutines`, and each goroutine will `add one record within time t`, where `t is the production time`.

Set the `delay time to m`.

Assume the `degree of concurrency is very low`, which is a conservative value.

Then the calculation formula for `delay time m = n*t`.

Assume the production time `t is very uniform and constant`.

The above formula becomes m proportional to t, meaning that `the delay time must be proportional to the number of goroutines n`.

However, how should t be measured?

If there is `OpenTelemetry`, it will be easy to count.

If the `spans can be combined`, it will be even `more convenient`.

Otherwise, unit tests will have to be written to estimate this value.

This value can definitely be found.

In short, I believe `the delay time should be proportional to the number of goroutines n`, and the program is also `easy to implement`.

(反正我认为延迟时间和协程数量成正比，其他数据不是给建议值，不然就用 OpenTelemetry 去统计)



## My opinion and suggestion

`OpenTelemetry` has a new `JSON-based` logging mechanism.

`MongoDB` is well suited to store these logs.

Currently, to use OpenTelemetry in Golang programs, many tools need to be connected, including visualization tools called `jaeger`, [third-party plugins](https://github.com/mongodb-labs/jaeger-mongodb), and `MongoDB`.

This is inconvenient and constrains `the popularity` of OpenTelemetry.

My suggestion is to allow Golang programs to `write OpenTelemetry logs directly to MongoDB`. Then we can query and analyze the logs in MongoDB directly, which is simpler than learning many different tools.

## More examples

Here are some golang examples to illustrate my view point:

```bash
$ pwd
# tracez

$ tree
.
├── [DIR] example
│   ├── [DIR] openTelemetry2file # <<<<< Example1: write OpenTelemetry logs directly to files
│   │   └── Makefile
│   └── [DIR] openTelemetry2mongodb # <<<<< Example2: write OpenTelemetry logs directly to MongoDB
│       └── Makefile
└── README.md
```

