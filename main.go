package main

import (
	"fmt"
	"github.com/mkideal/cli"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type argT struct {
	cli.Helper
	Dir    string `cli:"*dir" usage:"Project directory,eg:--dir=D:\hello"`
	Suffix string `cli:"suffix" usage:"The file extension,eg:--suffix=.java" dft:".java"`
	Crack  string `cli:"crack" usage:"Crack method.eg:--crack=d or --crack=r" dft:"d"`
}

func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		ctx.String("Dir=%s, Suffix=%s, Crack=%s\n", argv.Dir, argv.Suffix, argv.Crack)
		switch argv.Crack {
		case "d":
			Decryption(argv.Dir, argv.Suffix)
			break
		case "r":
			Rename(argv.Dir)
			break
		}
		return nil
	})

}
func Rename(Dir string) {
	files, err := WalkDir(Dir, ".hack")
	if err != nil {
		return
	}
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
		os.Rename(files[i], fmt.Sprintf("%s", files[i])[:len(files[i])-5])
	}
}

//解密
func Decryption(Dir, Suffix string) {
	files, err := WalkDir(Dir, Suffix)
	if err != nil {
		return
	}
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
		bStr, _ := ReadAll(files[i])
		ioutil.WriteFile(fmt.Sprintf("%s%s", files[i], ".hack"), bStr, 0666)
		os.Remove(files[i])
	}
}
func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix)                                                     //忽略后缀匹配的大小写
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}
func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
