# httpcat

`httpcat` is like UNIX's `cat` but for HTTP.

The main purpose of writing the util was to receive debug messages from a few REST endpoints simultaneously. When more than one source URLs are give it behaves rather like `tail -f`, prints messages from the sources by muxing the output line-by-line.

## Example

```shell
$ httpcat https://server{1..5}.lan/trace
```
