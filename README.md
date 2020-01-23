# homework

For all your self-assigned homework needs.

![](https://media.giphy.com/media/M7EQSsEXkGRvq/giphy.gif)

Inspired by [plutov/practice-go](https://github.com/plutov/practice-go).

## How to Use

### Test

From project root
```
go test ./<package_name>
```

### Benchmark

From project root
```
go test -run=xxx -bench=. ./<package_name>
```

### Examples

```
go test ./sparse_binary

go test -run=xxx -bench=. ./sparse_binary
```