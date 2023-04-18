# tracez

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

