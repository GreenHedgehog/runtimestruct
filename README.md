# runtimestruct
Generates Go-structs in runtime (with copied values) from human-readable format (JSON)

## Rules of JSON format transforming

```go
strct, err := NewFromJSON(io.Reader)
if err != nil {
    // handling errors of reading and decoding
}
fmt.Print("%+v\n", strct)
```

| JSON type | GO type         | comment |
| ---       | ---             | ---     |
| **Null**  | **nil**         | Would be typed nil of `*struct{}` type cause of `reflect` package limitations |
| **String**| **string**       | By `encoding/json` package |
| **Number**| **float64**      | By `encoding/json` package |
| **Array** | **[]interface{}**| By `encoding/json` package |
| **Object**| **struct{ ... }**| JSON field name would be copied to struct field. So, if that name starts with capital letter - it would be exported, otherwise - unexported. Fields would be in different order. |
