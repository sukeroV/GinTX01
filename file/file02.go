package main

import (
	"log"
	"os"
)

func main() {

	//file, err := os.Open("C:\\Users\\hsl\\Desktop\\text.txt")
	//if err != nil {
	//	fmt.Printf("open file err %v\n", err.Error())
	//	return
	//}

	data, err := os.ReadFile("D:\\mylog.log")
	if err != nil {
		log.Fatal(err)
	}
	
	os.WriteFile("D:\\mylog2.log", data, 0666)

	//print(os.Stdout.Write())

	//io.ReadFile("D:\\log_config.log")

	//defer func() {
	//	//关闭文件
	//	err := file.Close()
	//	if err != nil && file != nil {
	//		fmt.Printf("close file err %v", err.Error())
	//	}
	//}()
}
