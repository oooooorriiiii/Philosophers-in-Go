## Random

"think", "eat"の時間をランダムにしてみた。
どのような動作をするだろうか…

## 結果

![Result: random 20 Philos](static/Figure_1.png)

Philosopherごとに、食べる時間（eat）と考える時間（think）が異なりますが、
ちゃんと左右のリソース（fork）が使われていないときにしか食事（eat）ができていないことがわかりますね。

Mutexがちゃんと働いていることがわかります。