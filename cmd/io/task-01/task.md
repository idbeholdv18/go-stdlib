# 1. Реализуй свой Reader

Сделай структуру:

```go
type MyReader struct{}
```

которая реализует io.Reader и всегда возвращает строку "AAAAA...".

👉 цель:
понять контракт:

```go
Read(p []byte) (n int, err error)
```
