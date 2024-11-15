


## `analyze.py`の実行

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

## 参考文献

- [The Little Book of Semaphores](http://alumni.cs.ucr.edu/~kishore/papers/semaphores.pdf)
  -  P.93 ~
  -  AB Downey, 2005, University of California, Riverside
- [Parallel Programming (Carnegie Mellon University)](https://compeau.cbd.cmu.edu/wp-content/uploads/2016/08/lec_23.pdf)
- [Principles of Operating Systems (UC Irvine)](https://ics.uci.edu/~ardalan/courses/os/lectures/os_lecture_4.pdf)
