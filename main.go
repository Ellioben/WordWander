package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"handle-csv/handle"
	"io/ioutil"
	"log"
)

const (
	readDic   = "collection"
	outputDic = "output"
)

func main() {

	patchRow := []int{0, 1}

	fmt.Println("start--")
	files := searchFile(readDic)
	var allWord []string
	for _, file := range files {
		words := handle.ReadCsv(readDic, file, outputDic, patchRow)

		for _, word := range words {
			allWord = append(allWord, word)
		}
	}

	randomWord := handle.RandomWord(allWord)
	logrus.Infof("随机单词：%s", randomWord)

	//filePath, _ := fmt.Println("%s%s", path, fileName)
}

func searchFile(path string) []string {
	// 指定文件夹路径
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	var fileNames []string

	// 遍历文件列表并输出文件名称
	for _, file := range files {
		if !file.IsDir() {
			fileName := file.Name()
			fileNames = append(fileNames, fileName)
			logrus.Info(fileName)
		}
	}

	return fileNames

}
