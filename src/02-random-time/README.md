## Random

"think", "eat"の時間をランダムにしてみました。

```go
	thinkUnitTime := time.Duration(10 * rand.Intn(10))
	time.Sleep(time.Millisecond * thinkUnitTime)
```

```go
	eatUnitTime := time.Duration(10 * rand.Intn(10))
	time.Sleep(time.Millisecond * eatUnitTime)
```

どのような動作をするでしょうか…

## 結果

![Result: random 20 Philos](static/Figure_1.png)

Philosopherごとに、食べる時間（eat）と考える時間（think）が異なりますが、
ちゃんと左右のリソース（fork）が使われていないときにしか食事（eat）ができていないことがわかりますね。

[01-simple](../01-simple)で実行したときよりも、Mutexがちゃんと働いていることがわかりやすいですね。
