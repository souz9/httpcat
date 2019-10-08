# httpcat

`httpcat` is like UNIX's `cat` but for HTTP.

The main purpose of writing the util was to receive debug messages from a few REST endpoints simultaneously. When more than one source URLs are given it behaves rather like `tail -f`, prints messages from the sources by muxing the output line-by-line.

## How to Install

```shell
$ go get github.com/souz9/httpcat
```

## Usage

```shell
$ httpcat https://server{1..5}.lan/trace
```
