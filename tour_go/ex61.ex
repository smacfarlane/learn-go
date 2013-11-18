package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}
//Read(p []byte) (n int, err error)
func (r rot13Reader) Read(p []byte) (n int, err error) {
  n, err = r.r.Read(p)
    for i := range p {
        switch {
        case (p[i] >= 'a' && p[i] <= 'z') || (p[i] >= 'A' && p[i] <= 'Z'):
            p[i] += 13
            if(p[i] > 'Z' && p[i] < 'a') || p[i] > 'z' {
                p[i] -= 26
            }
        }
     }
    
  return
}

func main() {
    s := strings.NewReader(
        "Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
