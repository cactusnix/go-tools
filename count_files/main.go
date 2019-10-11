package main

import (
  "io/ioutil"
  log "github.com/sirupsen/logrus"
)

func main() {
  countFiles("/Users/windmill/documents/md")
}

func countFiles(filePath string) {
  count := 0
  extension := ""
  fileResult := make(map[string]int)
  files, err := ioutil.ReadDir(filePath)
  if err != nil {
    log.Fatal(err)
  }
  for _, file := range files {
    if file.IsDir() {
      log.Println(file.Name())
      countFiles(filePath + "/" + file.Name())
    } else {
      extension = ""
      count++
      for i := len(file.Name()) - 1; i >= 0; i-- {
        if string(file.Name()[i]) == "." {
          extension = file.Name()[i:]
          break
        }
      }
      fileResult[extension] = count
    }
  }
  log.Println(len(fileResult))
  log.Println(fileResult)
}
