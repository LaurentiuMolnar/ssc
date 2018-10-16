package main

import (
  "fmt"
  "github.com/tredoe/osutil/user/crypt/sha512_crypt"
)

func main() {
  c := sha512_crypt.New()
  hash, _ := c.Generate([]byte("<:cti18:>abpb"), []byte("$6$4LDtKhWO"))
  fmt.Println(hash)
}
