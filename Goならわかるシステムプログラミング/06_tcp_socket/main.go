package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

/**
基本的な構成
・サーバー：ソケットを開いて待ち受ける（net.listen）
・クライアント：ソケットに接続する（net.Dial）
*/

func main() {
	go setServer()
	time.Sleep(2 * time.Second)
	setClient()
}

// サーバーの起動
func setServer() {
	listener, err := net.Listen("tcp", "localhost:8888") // TCPソケットを開く
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")
	for { // 常にソケットへのリクエストを可能にするために無限ループ
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go func() { // スループットを上げるためにgoroutineで並列処理
			fmt.Printf("Accept %v\n", conn.RemoteAddr()) // 接続元を出力
			request, err := http.ReadRequest(
				bufio.NewReader(conn),
			)
			if err != nil {
				panic(nil)
			}
			dump, err := httputil.DumpRequest(request, true) // HTTPリクエストを丸ごと文字列化して返す
			if err != nil {
				panic(nil)
			}
			fmt.Println("server dump", string(dump))
			response := http.Response{ // レスポンスの作成
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       io.NopCloser(strings.NewReader("Hello World!\n")),
			}
			response.Write(conn) // レスポンスを返す
			conn.Close()
		}()
	}
}

// HTTPクライアント
func setClient() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest(
		"GET", "http://localhost:8888", nil, // Request Body が nil
	)
	if err != nil {
		panic(err)
	}
	request.Write(conn)
	response, err := http.ReadResponse(
		bufio.NewReader(conn), request,
	)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}
	fmt.Println("client dump", string(dump))
}
