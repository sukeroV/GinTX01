package main

//func main() {
//	//打开文件
//	file, err := os.Open("C:\\Users\\hsl\\Desktop\\text.txt")
//	if err != nil {
//		fmt.Printf("open file err %v\n", err.Error())
//		return
//	}
//
//	reader := bufio.NewReader(file)
//
//	for {
//		readString, err := reader.ReadString('\n') // 读取到换行就结束
//		if err == io.EOF {
//			break
//		}
//		fmt.Printf("%v\n", readString)
//	}
//
//	//看看文件是什么
//	fmt.Printf("file = %v", file)
//
//	defer func() {
//		//关闭文件
//		err := file.Close()
//		if err != nil && file != nil {
//			fmt.Printf("close file err %v", err.Error())
//		}
//	}()
//}
