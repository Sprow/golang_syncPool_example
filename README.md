Example usage of golang sync.Pool

`go test -v -bench=. -benchmem`

Benchmark text 1000 strings


```
BenchmarkWithPool1000
BenchmarkWithPool1000-4            13644             88147 ns/op          122941 B/op          1 allocs/op

BenchmarkWithoutPool1000
BenchmarkWithoutPool1000-4          4978            236430 ns/op          551877 B/op         11 allocs/op
```

Benchmark text 10000 strings

```
BenchmarkWithPool10000
BenchmarkWithPool10000-4            1072            941601 ns/op         1227706 B/op          6 allocs/op

BenchmarkWithoutPool10000
BenchmarkWithoutPool10000-4          368           3132069 ns/op         3919347 B/op         27 allocs/op```

