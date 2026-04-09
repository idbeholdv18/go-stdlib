# 5. MultiReader

Склей несколько readers:

```go
r1 := strings.NewReader("Hello ")
r2 := strings.NewReader("World")
```

👉 результат: "Hello World"

(можно сначала использовать io.MultiReader, потом написать свой)
