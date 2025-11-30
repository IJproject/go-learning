package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	S1()
	S2()
	S3()
}

// 標準入力
func S1() {
	buffer := make([]byte, 5)
	size, err := os.Stdin.Read(buffer) // 標準入力の内容をbufferに詰める
	if err == io.EOF {
		fmt.Println("EOF")
	}
	fmt.Printf("size=%d input='%s'\n", size, string(buffer))
}

// ファイル入力
func S2() {
	file, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}

// 初期化
func S3() {
	bBuffer := bytes.NewBufferString("bytes.NewBufferStringで初期化")
	bReader := strings.NewReader("strings.NewReaderで初期化")
	fmt.Println(bBuffer, bReader)
}

func S4() {

}
func S5() {

}
func S6() {

}
func S7() {

}
func S8() {

}
