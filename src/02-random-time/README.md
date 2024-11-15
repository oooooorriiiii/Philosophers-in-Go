## Random

"sleep", "eat"の時間をランダムにしてみました。

```go
	sleepUnitTime := time.Duration(10 * rand.Intn(10))
	time.Sleep(time.Millisecond * sleepUnitTime)
```

```go
	eatUnitTime := time.Duration(10 * rand.Intn(10))
	time.Sleep(time.Millisecond * eatUnitTime)
```

どのような動作をするでしょうか…

## 結果

![Result: random 20 Philos](static/Figure_1.png)

Philosopherごとに、食べる時間（eating）と睡眠時間（sleeping）、考える時間（Thinking）が異なりますが、
ちゃんと左右のリソース（fork）が使われていないときにしか食事（eat）ができていないことがわかりますね。

[01-simple](../01-simple)で実行したときよりも、Mutexがちゃんと働いていることがわかりやすいですね。

※）乱数なので、ところどころEating、Sleepingが0になるときが発生しているな。直すか？
