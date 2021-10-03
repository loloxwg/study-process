package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	dirName:="D:\\xwg\\Documents\\资料\\-DailyStudy"
	listFiles(dirName,0)
}


func listFiles(dirName string,level int){
	//level
	s:="|--"
	for i := 0; i < level; i++ {
		s="| "+s
	}
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fileInfos{
		fileName:=dirName+"/"+fi.Name()
		fmt.Printf("%s%s\n", s,fileName)
		if fi.IsDir() {
			listFiles(fileName,level+1)
		}

	}
}