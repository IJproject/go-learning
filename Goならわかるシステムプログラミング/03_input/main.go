package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	S1()
	S2()
	S3()
	S4()
	S5()
	S6()
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
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: example.com\r\n\r\n"))
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	fmt.Println(res.Header)
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}

func S5() {
	var source = "1行目\n2行目\n3行目"
	reader := bufio.NewReader(strings.NewReader(source))
	for { // 無限ループ
		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n", line)
		if err == io.EOF {
			break
		}
	}
}

func S6() {
	var source = "123 1.234 1.0e4 test"
	reader := strings.NewReader(source)
	var i int
	var f, g float64
	var s string
	fmt.Fscan(reader, &i, &f, &g, &s)
	fmt.Printf("i=%#v, f=%#v, g=%#v, s=%#v\n", i, f, g, s)
}
