# s

`s` command runs HTTP Server which serves static files on current directory.
`s` ってコマンドでそのディレクトリを document root としたhttpサーバが立ち上がる

## Installation インストール

```
go get github.com/hiroakis/s
```

or か

```
wget https://raw.githubusercontent.com/hiroakis/s/master/s_darwin_396
wget https://raw.githubusercontent.com/hiroakis/s/master/s_linux_386
wget https://raw.githubusercontent.com/hiroakis/s/master/s_windows_386.exe
```

## Run 実行

Just run `s`

```
s
```

## Option オプション

```
-h Listen IP address. Default: 0.0.0.0 リスンIP。デフォは"0.0.0.0"
-p Listen Port. Default: 8000 リスンポート。デフォは"8000"
-d The DocumentRoot. Default "." ドキュメントルートの指定。デフォは"."
```

## Motivation モチベーション

```
python -m SimpleHTTPServer
```

とか

```
python -m http.server 
```

とか

```
ruby -run -e httpd .
```

are pain for me and I forget sometime.
とかやるのすら面倒だったってのが。

## License

MIT 
