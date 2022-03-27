package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {
	banner()
	var (
		folderPath string
		oldSuffix  string
		newSuffix  string
	)
	var confirm string
	// 要求输入路径
	fmt.Println("输入要修改后缀的文件所在的路径(Enter the path of the file whose suffix you want to modify)：")
	_, _ = fmt.Scan(&folderPath)
	// 要求输入旧的后缀
	fmt.Println("输入你想更改的后缀类型(Enter the suffix you want to change)：")
	_, _ = fmt.Scan(&oldSuffix)
	// 要求输入新的后缀
	fmt.Println("你想把", oldSuffix, "改成什么？(What do you want to change", oldSuffix, "to ? )")
	_, _ = fmt.Scan(&newSuffix)
	// 确认
	fmt.Println("即将把[", folderPath, "]路径下的[", oldSuffix, "]类型文件更改为[", newSuffix, "]文件，确定吗？（Windows在完成后不会有提示）[Y/N]\n"+
		"About to change the type b file under the a path to the c file, are you sure? (Microsoft Windows won't prompt when done)[Y/N]")
	_, _ = fmt.Scan(&confirm)
	if confirm != "Y" {
		fmt.Println("并没有确定要更改，不做任何操作\n程序已经退出……\nNot sure to change, do nothing.\nExited...")
		return
	}
	// 路径处理，防止多余的分隔符，os.PathSeparator是当前操作系统下的路径分隔符
	if strings.HasSuffix(folderPath, string(os.PathSeparator)) {
		folderPath = folderPath[0 : utf8.RuneCountInString(folderPath)-1]
	}
	// 读取指定目录下的所有文件
	dir, err := ioutil.ReadDir(folderPath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("======================\n以下文件将受影响：\nThese files will be modify :")
	for _, files := range dir {
		// 判断是不是目标文件
		if strings.HasSuffix(files.Name(), "."+oldSuffix) && !files.IsDir() {
			fmt.Println(files.Name())
			split := strings.Split(files.Name(), "."+oldSuffix)[0] // 获取不包含文件扩展名的文件名
			// 重命名
			err := os.Rename(folderPath+string(os.PathSeparator)+split+"."+oldSuffix, folderPath+string(os.PathSeparator)+split+"."+newSuffix)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	fmt.Println("======================\n完成!done!")
	fmt.Println("按下enter退出……\nPress enter to exit……")
	b := make([]byte, 1)
	_, _ = os.Stdin.Read(b)
}

func banner() {
	fmt.Printf("%s",
		"███████╗    ██╗    ██╗         ███████╗\n"+
			"██╔════╝    ██║    ██║         ██╔════╝\n"+
			"█████╗      ██║    ██║         █████╗\n"+
			"██╔══╝      ██║    ██║         ██╔══╝\n"+
			"██║         ██║    ███████╗    ███████╗\n"+
			"╚═╝         ╚═╝    ╚══════╝    ╚══════╝\n\n")
	time.Sleep(time.Millisecond * 300)
	fmt.Printf("%s",
		"███████╗    ██╗   ██╗    ███████╗    ███████╗    ██╗    ██╗  ██╗\n"+
			"██╔════╝    ██║   ██║    ██╔════╝    ██╔════╝    ██║    ╚██╗██╔╝\n"+
			"███████╗    ██║   ██║    █████╗      █████╗      ██║     ╚███╔╝  \n"+
			"╚════██║    ██║   ██║    ██╔══╝      ██╔══╝      ██║     ██╔██╗  \n"+
			"███████║    ╚██████╔╝    ██║         ██║         ██║    ██╔╝ ██╗ \n"+
			"╚══════╝     ╚═════╝     ╚═╝         ╚═╝         ╚═╝    ╚═╝  ╚═╝ \n\n")
	time.Sleep(time.Millisecond * 300)
	fmt.Printf("%s",
		"███████╗    ██████╗     ██╗    ████████╗     ██████╗     ██████╗ \n"+
			"██╔════╝    ██╔══██╗    ██║    ╚══██╔══╝    ██╔═══██╗    ██╔══██╗\n"+
			"█████╗      ██║  ██║    ██║       ██║       ██║   ██║    ██████╔╝\n"+
			"██╔══╝      ██║  ██║    ██║       ██║       ██║   ██║    ██╔══██╗\n"+
			"███████╗    ██████╔╝    ██║       ██║       ╚██████╔╝    ██║  ██║\n"+
			"╚══════╝    ╚═════╝     ╚═╝       ╚═╝        ╚═════╝     ╚═╝  ╚═╝\n\n")
}
