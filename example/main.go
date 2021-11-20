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

  fmt.Println(s)

  // Length
  fmt.Println(lsmk.Length())

  // Get
  v, err := lsmk.Get("testkey")

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(v)

  // Set 2
  s2, err := lsmk.Set("testkey2", "testvalue2")
 
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(s2)

  // Length
  fmt.Println(lsmk.Length())

  // Get 2
  v2, err := lsmk.Get("testkey2")

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(v2)

  // Remove
  s3, err := lsmk.Remove("testkey")

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(s3)

  // Length
  fmt.Println(lsmk.Length())

  // Get 3
  v3, err := lsmk.Get("testkey")

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(v3)
}

