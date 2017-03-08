// https://golang.org/pkg/crypto/
package main

import (
  "fmt"
  "crypto/md5"
  "crypto/sha1"
  "crypto/sha256"
  "crypto/sha512"
  "hash"
)


func main() {
  funcNameList := []string{"MD5", "SHA1", "SHA224", "SHA256", "SHA384", "SHA512"}
  funcMap := map[string]func(msg []byte) string{
    "MD5"     : func(msg []byte) string{var h hash.Hash = md5.New();h.Write(msg);return string(h.Sum(nil))},
    "SHA1"    : func(msg []byte) string{var h hash.Hash = sha1.New();h.Write(msg);return string(h.Sum(nil))},
    "SHA224"  : func(msg []byte) string{var h hash.Hash = sha256.New224();h.Write(msg);return string(h.Sum(nil))},
    "SHA256"  : func(msg []byte) string{var h hash.Hash = sha256.New();h.Write(msg);return string(h.Sum(nil))},
    "SHA384"  : func(msg []byte) string{var h hash.Hash = sha512.New384();h.Write(msg);return string(h.Sum(nil))},
    "SHA512"  : func(msg []byte) string{var h hash.Hash = sha512.New();h.Write(msg);return string(h.Sum(nil))},
  }
  var msg1 string
  fmt.Printf("Input string : ")
  fmt.Scanf("%s", &msg1)
  for _,funcName := range funcNameList{
    fmt.Printf("%s \t: %x\n", funcName, funcMap[funcName]([]byte(msg1)))
  }
}
