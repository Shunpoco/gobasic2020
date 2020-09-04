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
BenchmarkPopCountLoop-4         70181626                17.2 ns/op
BenchmarkPopCountUnit-4         1000000000               0.622 ns/op
PASS
ok      pop/popcount    1.968s
```