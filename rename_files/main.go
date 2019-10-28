package main

import (
  "io/ioutil"
  log "github.com/sirupsen/logrus"
  "os"
  "strconv"
  "fmt"
)

func main() {
  var funcNum string
  fmt.Println("输入需要使用功能的序号：")
  fmt.Println("1. 批量重命名文件")
  fmt.Println("2. 替换文件夹部分名称")
  fmt.Scanf("%s", &funcNum)
  invokeFunc(funcNum)
}
func invokeFunc(funcNum string) {
  if len(funcNum) == 1 {
    switch funcNum {
    case "1":
      renameFile()
    case "2":
      replaceStr()
    }
  } else {
    fmt.Println("请输入正确的功能序号")
  }
}
func replaceStr() {
  var oldStr string
  var newStr string
  var filePath string
  fmt.Println("输入需要替换的旧字符串")
  fmt.Scanf("%s", &oldStr)
  fmt.Println("输入需要替换的新字符串")
  fmt.Scanf("%s", &newStr)
  fmt.Println("输入需要批量修改文件的路径")
  fmt.Scanf("%s", &filePath)
  files, err := ioutil.ReadDir(filePath)
  if err != nil {
    log.Println("读取路径错误：", err)
  }
  for _, file := range files {
    if file.IsDir() {
      oldName := file.Name()
      newName := ""
      for _, str := range oldName {
        if string(str) == "_" {
          newName += "-"
        } else {
          newName += string(str)
        }
      }
      err := os.Rename(filePath + oldName, filePath + newName)
      if err != nil {
          log.Fatal("reName Error", err)
      } else {
        log.Info("替换成功！")
      }
    } else {
      continue
    }
  }
}

func renameFile() {
  var filePath string
  var fileName string
  start := 1
  fmt.Println("输入需要批量修改文件的路径")
  fmt.Scanf("%s", &filePath)
  fmt.Println("输入修改的新名字")
  fmt.Scanf("%s", &fileName)
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
