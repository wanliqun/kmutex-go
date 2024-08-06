# kmutex-go

`kmutex-go` is an efficient key-based multiple mutex library for Go, designed to improve locking performance by preventing a giant time-consuming locking. This library is ideal for scenarios where you need to lock based on a specific key, avoiding contention on a single mutex.

## Features

- Key-based locking to reduce contention
- Efficient memory usage with sync.Pool
- Easy-to-use API

## Installation

```sh
go get github.com/wanliqun/kmutex-go
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/wanliqun/kmutex-go"
)

func main() {
    key := "example_key"
    kmutex.Lock(key)
    // Perform your operations here
    fmt.Println("Locked by key:", key)
    kmutex.Unlock(key)
    fmt.Println("Unlocked by key:", key)
}
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
