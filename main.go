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

	var allWord []string

	for _, file := range files {

		// todo 先获取到当前file的单词，写入新文件
		words := handle.GetWordFromFileAndWrite2CSV(readDic, file, outputDic, patchRow, allWord)
		continueinfo := declareUpdate(file)

		if continueinfo {
			continue
		}
		// todo 当前单词和存起来的所有单词差集，
		// 单词差级
		words = handle.UniqueResource(words, allWord)
		// todo 文件过滤掉后，单词也需要过滤掉（单词还是需要读的，读完后过滤然后再生成文章）

		logrus.Infof("文件名：《%s》，单词组：%v，一共：%v 个\n", file, words, len(words))

		appendNewArticle(words, file)
		allWord = append(allWord, words...)

	}

	if len(allWord) == 0 {
		logrus.Infof("没有新文件，请检查colleation文件夹")

	} else {

		logrus.Infof("一共%v\n", len(allWord))
		randomWord := handle.RandomWord(allWord)
		logrus.Infof("随机单词：%s", randomWord)
	}

	//filePath, _ := fmt.Println("%s%s", path, fileName)
}

func declareUpdate(file string) bool {
	// todo 再获取已有的文件
	existFiles := searchFile(outputDic)
	// todo 选择差集

	continueinfo := false

	// 文件差
	for _, existfile := range existFiles {
		if file == existfile {
			continueinfo = true
		}
	}
	return continueinfo
}

func appendNewArticle(words []string, file string) {
	article := GeneratorArticle(words)

	//todo 每个文章生成对应日期的文件夹里
	handle.WriteArticle2CSV(outputDic, file, article)
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
