package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"handle-csv/handle"
	"io/ioutil"
	"log"
	"os/exec"
)

const (
	readDic    = "collection"
	outputDic  = "output"
	scriptPath = "./generator-article"
)

func main() {

	patchRow := []int{0, 1}

	// todo 三个月调用relingo生成一次文件

	fmt.Println("start--")
	files := searchFile(readDic)
	var allWord []string
	for _, file := range files {
		words := handle.ReadCsv(readDic, file, outputDic, patchRow)

		words = uniqueWords(words, allWord)
		logrus.Infof("文件名：《%s》，单词组：%v，一共：%v 个\n", file, words, len(words))
		//todo 余下来的单词调用ai的api，生成文章
		GeneratorArticle()

		for _, word := range words {
			allWord = append(allWord, word)
		}

		//todo 每个文章生成对应日期的文件夹里

	}

	logrus.Infof("一共%v\n", len(allWord))
	randomWord := handle.RandomWord(allWord)
	logrus.Infof("随机单词：%s", randomWord)

	//filePath, _ := fmt.Println("%s%s", path, fileName)
}

func GeneratorArticle() {
	// 定义 Python 脚本路径和参数
	arg1 := "xxxx"

	// 构建命令行参数
	args := []string{scriptPath, arg1}

	// 执行 Python 脚本
	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	// 输出 Python 脚本的执行结果
	fmt.Println(string(output))
}

func uniqueWords(words []string, allWord []string) []string {
	var uniqueWords []string

	for _, word := range words {
		found := false
		for _, s := range allWord {
			if word == s {
				found = true
				break
			}
		}
		if !found {
			uniqueWords = append(uniqueWords, word)
		}
	}

	words = uniqueWords
	return words
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
		}
	}

	return fileNames

}
