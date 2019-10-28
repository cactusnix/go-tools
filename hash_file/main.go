package main

import (
	"log"
	"encoding/hex"
  "crypto/sha256"
)

func main()  {
  caculateHash("test the hash")
  caculateHash("test thehash")
}

func caculateHash(toBeHashed string) {
  // 这个命名变量的方式优秀啊！含义和类型都有了。
  hashInBytes := sha256.Sum256([]byte(toBeHashed))
  hashInStr := hex.EncodeToString(hashInBytes[:])
  log.Printf("%s, %s", toBeHashed, hashInStr)
}
