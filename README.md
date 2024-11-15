# Dining philosophers problem

> In computer science, the dining philosophers problem is an example problem often used in concurrent algorithm design to illustrate synchronization issues and techniques for resolving them. - Wikipedia

![Dining_philosophers_diagram](https://github.com/user-attachments/assets/9f43aaad-fcd4-4d3a-a9c8-615b05762d2f)



## 参考文献

- [The Little Book of Semaphores](http://alumni.cs.ucr.edu/~kishore/papers/semaphores.pdf)
  -  P.93 ~
  -  AB Downey, 2005, University of California, Riverside
- [Parallel Programming (Carnegie Mellon University)](https://compeau.cbd.cmu.edu/wp-content/uploads/2016/08/lec_23.pdf)
- [Principles of Operating Systems (UC Irvine)](https://ics.uci.edu/~ardalan/courses/os/lectures/os_lecture_4.pdf)

## その他

### `main.go`の実行

`main.go`のあるディレクトリで以下を実行

```bash
$ go run .
```

analyze.pyにいれるlogファイルの出力

```bash
$ go run . > actions.log # ファイル名はactions.logで固定です
```

### `analyze.py`の実行

必要ライブラリ

```
pandas
matplotlib
```

```bash
# 環境の新規構築(最初の1回だけ)
$ python3 -m env venv

# 環境の有効化（検証環境に入るたびに実行）
$ source venv/bin/activate

# ライブラリのインストール (最初の1回だけ)
$ pip install matplotlib pandas

# プログラムの実行（検証環境に入るたびに実行）
$ python3 analyze.py

# 環境から抜ける（検証環境に入るたびに実行）
$ deactivate
```
