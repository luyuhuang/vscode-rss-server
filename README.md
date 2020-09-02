# vscode-rss-server

Just for Inoreader OAuth. The main project See [https://github.com/luyuhuang/vscode-rss](https://github.com/luyuhuang/vscode-rss).

## Build & Usage

This is a simple HTTP server write by Go. It's easy to build and run.

```sh
git clone https://github.com/luyuhuang/vscode-rss-server.git
cd vscode-rss-server
go build
./vscode-rss-server
```

Or using `go get`:

```sh
go get github.com/luyuhuang/vscode-rss-server
vscode-rss-server
```

It'll listen `127.0.0.1:8080` by default. You can specify it in the first argument. For example:

```sh
vscode-rss-server 0.0.0.0:80
```
