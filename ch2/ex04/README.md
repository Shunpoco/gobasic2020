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
BenchmarkPopCountBitShift-4     22253760                58.1 ns/op
BenchmarkPopCountUnit-4         1000000000               0.326 ns/op
PASS
ok      pop/popcount    1.784s```