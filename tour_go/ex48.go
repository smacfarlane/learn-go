package main

import (
    "fmt"
  "math/cmplx"
)

func Cbrt(x complex128) complex128 {
  var z complex128 = 1.0

  for i := 0; i < 10; i++ {
      z = z - ((cmplx.Pow(z,3.0) - x) / (3 * z * z))
    }
  return z
}

func main() {
    fmt.Println(Cbrt(2))
}
