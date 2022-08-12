[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/kaatinga/robinson/blob/main/LICENSE)
[![lint workflow](https://github.com/kaatinga/robinson/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/kaatinga/robinson/actions?query=workflow%3Alinter)
[![help wanted](https://img.shields.io/badge/Help%20wanted-True-yellow.svg)](https://github.com/robinson/strconv/issues?q=is%3Aopen+is%3Aissue+label%3A%22help+wanted%22)

# robinson

A single value thread-safe cache for any value type. 

Minimalistic implementation is used, based on Go generics. Therefore the cache is strongly typed.
Once created, the cache allows to set only the type that was used to create cache.

The cache is error-free.

```go
value := 123
crusoe := NewCrusoe[int]()
crusoe.Set(value)
cacheValue := crusoe.Get()
```