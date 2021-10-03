package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//fileName:="D:\\xwg\\Documents\\资料\\-DailyStudy\\BBB.txt"
	//data, err := ioutil.ReadFile(fileName)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(data)
	//fmt.Println(string(data))

	//s1:="hello world aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	//
	//err := ioutil.WriteFile(fileName, []byte(s1), os.ModePerm)
	//fmt.Println(err)

	//dirName:="D:\\xwg\\Documents\\资料\\-DailyStudy"
	//fileInfos,err:=ioutil.ReadDir(dirName)
	//if err!=nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(len(fileInfos))
	//for i := 0; i < len(fileInfos); i++ {
	//	fmt.Printf("第%d个:名称:%s 是否是目录: %t\n",i,fileInfos[i].Name(),fileInfos[i].IsDir())
	//}

	dir, err := ioutil.TempDir("D:\\xwg\\Documents\\资料\\-DailyStudy", "Test")
	if err!=nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(dir)
	fmt.Println(dir)
	file, err := ioutil.TempFile(dir, "test")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file)

	defer os.RemoveAll(file.Name())

}
