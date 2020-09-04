## 実行方法
- `./run.sh` を実行
## 検証サイト
- `https://www.youtube.com`
## 検証結果
```sh
$ ./run.sh 
+ : Set Variables
+ readonly SITES=https://www.youtube.com/
+ SITES=https://www.youtube.com/
+ : Run go code
+ go run main.go https://www.youtube.com/
0.44s  419520 https://www.youtube.com/
0.23s  509602 https://www.youtube.com/
0.67s elapsed
```
- 以上の結果から、何かしらのキャッシュが行われており2度目のfetchでは時間が短縮されているのがわかる
- バイト数も異なるので、異なるコンテンツを取得していると推察されるが、具体的に何が異なるかまでは検証できず