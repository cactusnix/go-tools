package main

import (
  "io/ioutil"
  log "github.com/sirupsen/logrus"
  "os"
  "strconv"
  "fmt"
)

func main() {
  var filePath string
  var fileName string
  fmt.Println("输入需要使用功能的序号：")
  fmt.Println("1. 批量重复名文件")
  fmt.Println("输入需要使用功能的序号：")
  fmt.Println("输入需要批量修改文件的路径")
  fmt.Scanf("%s", &filePath)
  fmt.Println("输入修改的新名字")
  fmt.Scanf("%s", &fileName)
  renameFile(filePath, fileName, 1)
}

func replaceStr(newStr string)  {

}

func renameFile(filePath string, fileName string, start int) {
  files, err := ioutil.ReadDir(filePath)
  if err != nil {
    log.Println(err)
  }
  log.Println(len(files))
  for index, file := range files {
    if file.IsDir() {
      log.Fatal("这是一个文件夹，暂不支持迭代遍历，危险性较高！")
    } else {
      oldName := file.Name()
      ext := ""
      for i := len(file.Name()) - 1; i >= 0; i-- {
        if string(file.Name()[i]) == "." {
          ext = file.Name()[i:]
          break
        }
      }
      log.Println(ext, index)
      newName := fileName + "-" + strconv.Itoa(index + start) + ext
      log.Println(oldName)
      log.Println(filePath + newName)
      if isDuplicate(filePath + newName) {
        log.Fatal("重命名文件重复，会覆盖原文件，请更换名字。")
      } else {
        err := os.Rename(filePath + oldName, filePath + newName)
        if err != nil {
            log.Fatal("reName Error", err)
        }
      }
    }
  }
}
func isDuplicate(filePath string) bool {
  _, err := os.Stat(filePath)
  if err != nil {
    if os.IsExist(err) {
      return true
    }
    return false
  }
  return true
}
