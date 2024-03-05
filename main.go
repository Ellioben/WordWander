package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"handle-csv/handle"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
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
	existFile := searchFile(outputDic)

	var allWord []string
	handFile := uniqueResource(files, existFile)

	for _, file := range handFile {

		words := handle.ReadCsv(readDic, file, outputDic, patchRow)

		words = uniqueResource(words, allWord)
		// todo 文件过滤掉后，单词也需要过滤掉（单词还是需要读的，读完后过滤然后再生成文章）
		logrus.Infof("文件名：《%s》，单词组：%v，一共：%v 个\n", file, words, len(words))
		//words := handle.ReadCsv(readDic, file, outputDic, patchRow)

		article := GeneratorArticle(words)

		for _, word := range words {
			allWord = append(allWord, word)
		}

		//todo 每个文章生成对应日期的文件夹里
		handle.WriteArticle2CSV(outputDic, file, article)
	}

	logrus.Infof("一共%v\n", len(allWord))
	randomWord := handle.RandomWord(allWord)
	logrus.Infof("随机单词：%s", randomWord)

	//filePath, _ := fmt.Println("%s%s", path, fileName)
}

func GeneratorArticle(arg1 []string) string {
	arg1String := strings.Join(arg1, ",")

	args := []string{scriptPath, arg1String}

	// 执行 脚本
	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	// 输出 Python 脚本的执行结果
	return string(output)

}

func uniqueResource(collection []string, allCollection []string) []string {
	var uniqueWords []string

	for _, x := range collection {
		found := false
		for _, s := range allCollection {
			if x == s {
				found = true
				break
			}
		}
		if !found {
			uniqueWords = append(uniqueWords, x)
		}
	}

	collection = uniqueWords
	return collection
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
