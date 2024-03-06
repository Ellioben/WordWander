package handle

import (
	"encoding/csv"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

func ReadCsv(path string, fileName string, outputDic string, content []string) []string {
	var str = []string{path, fileName}
	filePath := strings.Join(str, "/")

	// 打开文件(只读模式)，创建io.read接口实例
	opencast, err := os.Open(filePath)
	if err != nil {
		logrus.Error("csv文件打开失败！")
	}
	defer opencast.Close()
	newFilePath := ""
	// 创建csv读取接口实例
	reader := csv.NewReader(opencast)
	var lines []string

	newfile := []string{outputDic, fileName}
	newFilePath = strings.Join(newfile, "/")

	line, err := reader.Read()
	if err == io.EOF {
		return []string{}
	} else if err != nil {
		fmt.Println("读取文件Error: ", err)
		return []string{}
	}

	// 输出提取的数据到新文件
	if line[0] != "" {

		WriterCSV(newFilePath, content)

	}

	return lines
}

func GetWordFromFile(path string, fileName string, row []int) []string {
	var str = []string{path, fileName}
	filePath := strings.Join(str, "/")

	// 打开文件(只读模式)，创建io.read接口实例
	opencast, err := os.Open(filePath)
	if err != nil {
		logrus.Error("csv文件打开失败！")
	}
	defer opencast.Close()
	// 创建csv读取接口实例
	reader := csv.NewReader(opencast)
	var lines []string

	for {

		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("读取文件Error: ", err)
			return []string{}
		}

		// 输出提取的数据到新文件
		if line[0] != "" {
			lines = append(lines, line[row[0]])
			var goalString []string

			for i := 0; i < len(row); i++ {

				// catch the word to print

				// generator newfile
				goalString = append(goalString, line[row[i]])
				// todo 这里需要one by one 写入

			}
		}
	}
	return lines
}

func addLineBreak() string {
	s := []string{"\n"}
	s = append(s, "====================================================")
	s = append(s, "\n")
	line := "\n" + strings.Join(s, "") + "\n"
	return line
}

func WriterCSV(path string, str []string) {

	//OpenFile读取文件，不存在时则创建，使用追加模式
	File, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		logrus.Error("文件打开失败！")
	}
	defer File.Close()

	//创建写入接口
	WriterCsv := csv.NewWriter(File)
	//str:=[]string{"chen1","hai1","wei1"} //需要写入csv的数据，切片类型

	//写入一条数据，传入数据为切片(追加模式)
	err1 := WriterCsv.Write(str)
	if err1 != nil {
		fmt.Println("WriterCsv写入文件失败")
	}
	WriterCsv.Flush() //刷新，不刷新是无法写入的
	logrus.Debug("数据写入成功...")
}

func WriterEndCSV(path string, str string) {

	//OpenFile读取文件，不存在时则创建，使用追加模式

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrus.Error("无法打开文件：%v", err)
	}
	defer file.Close()

	_, err = file.WriteString(str)
	if err != nil {
		logrus.Error("写入文章失败：%v", err)
	}
}

func WriteArticle2CSV(path, fileName, content string) {
	str := []string{path, fileName}
	newFilePath := strings.Join(str, "/")
	logrus.Infof("生成新文章--->%s", newFilePath)

	file, err := os.OpenFile(newFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrus.Error("无法打开文件：%v", err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		logrus.Error("写入文章失败：%v", err)
	}
}
