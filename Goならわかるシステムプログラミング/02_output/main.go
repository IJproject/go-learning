package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

// File Descriptor: カーネルレイヤーで用意されている、アクセスできる対象に割り振られたIDのようなもの（ファイルやソケット、OSなど）
// io.Writerは、OSごとの差異を吸収する抽象化レイヤーとしての役割を持つ
func main() {
	S1()
	S2()
	S3()
	S4()
	S5()
}

// ファイル出力
func S1() {
	file, err := os.Create("./outputs/01_test.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("os.File Example\n"))
	file.Close()
}

// 画面出力
func S2() {
	// fmt.Printlnでは、Fileの構造体として「os.Stdout」を渡しているため、コマンドラインに出力されるようになっている
	os.Stdout.Write([]byte("os.Stdout Example\n"))
}

// バッファ（Goアプリ内のオンメモリキャッシュで読み書き可能）
func S3() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer Example")) // オンメモリにバイトデータを入れておく
	fmt.Println(buffer.String())
}

// バッファ（文字列連結特化の読み取り専用バッファ）
func S4() {
	var builder strings.Builder
	builder.Write([]byte("strings"))
	builder.Write([]byte(".Builder"))
	builder.Write([]byte(" Example"))
	fmt.Println(builder.String())
}

// インターネットアクセスの送信
// httpパッケージを使用すれば、もっと簡潔に表現することはできる
func S5() {
	conn, err := net.Dial("tcp", "example.com:80") // 指定したネットワーク宛に接続して、通信路を開く関数
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: example.com\r\n\r\n") // ソケットに書き込むことでHTTPリクエストを発火
	io.Copy(os.Stdout, conn)
}

func S6() {

}

func S7() {

}
