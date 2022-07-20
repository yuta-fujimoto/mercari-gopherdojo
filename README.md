# mercari-gopherdojo-00

# 概要

画像の形式を変換するコマンド

# buid
```
go mod tidy
go build
```

# usage

```
./convert -i=[input format] -o=[output format] [image file path]
[input format]はpng, jpeg, gifに対応
[output format]はpng, jpeg, gif, pgm, ppmに対応
```
ppm, pgmはPNMの一種です
https://ja.wikipedia.org/wiki/PNM_(%E7%94%BB%E5%83%8F%E3%83%95%E3%82%A9%E3%83%BC%E3%83%9E%E3%83%83%E3%83%88)
pgmに変換する際には画像が白黒になります。

# test
```
bash ./exec_test.sh
```

ただし、mercari-gopherdojo-01でテストもgoで書くように変更しました。
