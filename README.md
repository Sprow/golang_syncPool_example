Example usage of golang sync.Pool

`go test -v -bench=. -benchmem`

Benchmark text 1000 strings


```
BenchmarkWithPool1000
BenchmarkWithPool1000-4            13422             90092 ns/op          122896 B/op          1 allocs/op
BenchmarkWithoutPool1000
BenchmarkWithoutPool1000-4          4897            249630 ns/op          551840 B/op         11 allocs/op
```

Benchmark text 10000 strings

```
BenchmarkWithPool10000
BenchmarkWithPool10000-4            1152           1146314 ns/op         1223172 B/op          1 allocs/op
BenchmarkWithoutPool10000
BenchmarkWithoutPool10000-4          348           3395771 ns/op         3914023 B/op         14 allocs/op
```