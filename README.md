[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/kaatinga/robinson/blob/main/LICENSE)
[![lint workflow](https://github.com/kaatinga/robinson/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/kaatinga/robinson/actions?query=workflow%3Alinter)
[![help wanted](https://img.shields.io/badge/Help%20wanted-True-yellow.svg)](https://github.com/robinson/strconv/issues?q=is%3Aopen+is%3Aissue+label%3A%22help+wanted%22)

# Robinson

The "Robinson" package is a lightweight and efficient Go library that provides a thread-safe cache implementation for storing and retrieving values of any type. It offers a simple and minimalistic approach to caching, leveraging the power of Go generics to ensure strong typing and optimal performance.

Key Features:

- Thread-Safe Caching: The Robinson package offers a thread-safe cache that allows concurrent access from multiple goroutines without the risk of data corruption or race conditions. It provides synchronized operations to ensure safe and reliable caching in concurrent environments.
- Strong Typing: With Go generics, the cache enforces strong typing, meaning that once the cache is created, it can only store and retrieve values of the exact type specified during initialization. This ensures type safety and prevents accidental data mismatches or type-related errors.
- Error-Free Operation: The Robinson library is designed to provide a seamless and error-free caching experience. It handles internal operations and resource management efficiently, reducing the possibility of unexpected errors or exceptions during cache operations.
- Zero Dependencies: The package has been developed to be lightweight and self-contained, with no external dependencies. It minimizes the overall package size and simplifies integration into your projects, avoiding potential conflicts or compatibility issues with other libraries.

```go
value := 123

// Create a cache for storing integer value
crusoe := NewCrusoe[int]()

// Set the value in the cache
crusoe.Set(value)

// Get the value from the cache
cacheValue := crusoe.Get()
```

Whether you need to implement a simple in-memory cache for your application or require a thread-safe caching solution for concurrent operations, the Robinson package provides a reliable and efficient caching mechanism. Its minimalistic design, strong typing, and error-free operation make it a valuable tool for improving performance and enhancing data management in your Go projects.