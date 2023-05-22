package file

import (
	"fmt"
	"log"
	"os"
)

func Search() {
	// 開啟文件
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	// 判斷文件是否被編輯
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	// 讀取文件的大小
	fileSize := fileInfo.Size()
	// 創建一個slice以alloc足夠的空間
	fileBytes := make([]byte, fileSize)

	// 讀取文件到slice
	_, err = file.Read(fileBytes)
	if err != nil {
		log.Fatal(err)
	}

	// 打印文件內容
	fmt.Println("string", string(fileBytes))

	// 關閉文件
	file.Close()
}
