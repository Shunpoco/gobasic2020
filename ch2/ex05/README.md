## Benchmark実行
```sh
cd ./popcount
go test -bench=.
```
## 結果
```sh
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: pop/popcount
BenchmarkPopCountExclusive-4    51482412                25.0 ns/op
BenchmarkPopCountUnit-4         1000000000               0.661 ns/op
PASS
ok      pop/popcount    2.122s```