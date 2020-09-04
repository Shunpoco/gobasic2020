## Benchmark実行
- `go test -bench=.`
## 結果
```sh
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: ex03
BenchmarkSlowEcho-4           31          37459864 ns/op
BenchmarkFastEcho-4        10000            108175 ns/op
PASS
ok      ex03    2.376s
```
- strings.Join()を行ったコードの方が明らかに早い