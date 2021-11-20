# lsm

This is a Go implementation of like a **l**ocal**S**torage **m**ock.

However, please note that it is arranged to be implemented in Go, and does not fully reproduce the behavior of localStorage.

## Usage

```go
package main

import (
  "fmt"
  "log"
  "github.com/shinshin86/lsm"
)

func main() {
  lsmk := lsm.New()

  // Set
  s, err := lsmk.Set("testkey", "testvalue")
 
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(s) // map[testkey:testvalue]

  // Length
  fmt.Println(lsmk.Length()) // 1

  // Get
  v, err := lsmk.Get("testkey")

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(v) // testvalue

  // Set 2
  s2, err := lsmk.Set("testkey2", "testvalue2")
 
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(s2) // map[testkey:testvalue testkey2:testvalue2]

  // Length
  fmt.Println(lsmk.Length()) // 2

  // Get 2
  v2, err := lsmk.Get("testkey2")

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(v2) // testvalue2

  // Remove
  s3, err := lsmk.Remove("testkey")

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(s3) // map[testkey2:testvalue2]

  // Length
  fmt.Println(lsmk.Length()) // 1

  // Get 3
  v3, err := lsmk.Get("testkey")

  if err != nil {
    log.Fatal(err) // A non-existent key has been specified.
  }

  fmt.Println(v3) // This code will not be executed
}
```

## Licence

[MIT](https://github.com/shinshin86/go-lsm/blob/main/LICENSE)

## Author

[Yuki Shindo](https://shinshin86.com/en)

