package handle

import (
	"encoding/csv"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

func ReadCsv(path string, fileName string, outputDic string, row []int) []string {
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
		str := []string{outputDic, fileName}
		newFilePath := strings.Join(str, "/")

		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("读取文件Error: ", err)
			return []string{}
		}

		// 输出提取的数据到新文件
		if line[0] != "" {
			var goalString []string
			lines = append(lines, line[row[0]])

			for i := 0; i < len(row); i++ {

				// catch the word to print

				// generator newfile
				goalString = append(goalString, line[row[i]])

			}

			WriterCSV(newFilePath, goalString)

			// todo 写入文章

		}
	}

	return lines
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
