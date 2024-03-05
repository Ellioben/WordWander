package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// 定义 Python 脚本路径和参数
	scriptPath := "generator-article"
	arg1 := "art"

	// 构建命令行参数
	chmod := []string{"chmod +x ../"}
	args := []string{"../", scriptPath, arg1}

	exec.Command(chmod[0], args[1:]...)

	// 执行 Python 脚本
	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	// 输出 Python 脚本的执行结果
	fmt.Println(string(output))
}
