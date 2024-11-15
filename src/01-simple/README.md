## デッドロック防止-1

- 課題）哲学者が同時にフォークを取ろうとすると、デッドロック（全員が待ち状態になり、進まなくなる）が発生する可能性があります。
- 解決方法）哲学者が左右どちらのフォークを先に取るのかを制御する
  - 奇数番号の哲学者は右→左の順番にフォークを取る
  - 偶数番号の哲学者は左→右の順番にフォークを取る

## 利用機能

- `sync.Mutex`
  - 排他制御を行なうためにリソースをロックするための機能
  - 1つの処理だけがあるリソースにアクセスできるようにし、他の処理が同時にアクセスしてしまうことを防ぐ
  - 本実装
    - リソース: `fork[i]`
      - `philosopher`が`eat()`するために必要な`fork`。
      - `eat()`するためには`fork`が2本必要。
    - 処理: `philosopher.eat()`
- `sync.WaitGroup`
  - 複数のgoroutineが終了するまで待機するためのカウンタ
```go
func (p *Philosopher) dine(wg *sync.WaitGroup) {
    // deferを利用して、この関数が終了したときに呼ばれるようにする。
	// 終了するごとにカウントを1減らす
	defer wg.Done()
        ...
}

...

func main () {
	...
	for i := 0; i < numPhilosophers; i++ {
		...
		wg.Add(1) // 処理の数（哲学者の数、Threadの数）分だけカウントする
		go philosophers[i].dine(&wg)
	}
}
```

## 結果

![Result: deadlock prevention 5 Philos](static/Figure_1.png)

Philosopherは左右のリソース（Fork）が使われていないときにしか食事（"eat"）ができていないことがわかりますね。