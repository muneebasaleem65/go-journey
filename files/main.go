package main

import (
	// "bufio"
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("File Name: ", fileInfo.Name())
	// fmt.Println("File or Folder: ", fileInfo.IsDir())
	// fmt.Println("File Size: ", fileInfo.Size())
	// fmt.Println("File Permission: ", fileInfo.Mode())
	// fmt.Println("File Modified at: ", fileInfo.ModTime())

	//read file

	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// buf := make([]byte, 12)

	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }

	// for i:=0; i<len(buf); i++{
	// 	fmt.Println("data", d, string(buf[i]))
	// }

	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))


	//read directory

	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfos, err := dir.ReadDir(-1)

	// for _, fi := range fileInfos {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	//create file

	// f, err:=os.Create("example2.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// f.WriteString("Hi Go")

	//read and write to another file

	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)

	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }
	// writer.Flush()

	// fmt.Println("Written to a new file successfully")

	//delete file

	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("File deleted")
}
