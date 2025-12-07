package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// go setServer()
	// time.Sleep(2 * time.Second)
	// setClient()

	go setMultiCastServer()
	setMultiCastClient()
}

func setServer() {
	fmt.Println("Set up Server")
	conn, err := net.ListenPacket("udp", "localhost:8889")
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	buffer := make([]byte, 1500) // パケットの受け皿になるバッファ
	for {
		length, remoteAddress, err := conn.ReadFrom(buffer) // bufferに受信データを書き込む
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received from %v: %v\n", remoteAddress, string(buffer[:length]))
		_, err = conn.WriteTo([]byte("Hello from Server"), remoteAddress)
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
		_, err = conn.WriteTo([]byte("Hello from Server 2"), remoteAddress)
		if err != nil {
			panic(err)
		}
	}
}

func setClient() {
	conn, err := net.Dial("udp4", "localhost:8889")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Sending to Server")
	_, err = conn.Write([]byte("Hello from Client"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Receiving from Server")
	buffer := make([]byte, 1500)
	for {
		_ = conn.SetReadDeadline(time.Now().Add(2 * time.Second)) // 2秒以上レスポンスがなければタイムアウト
		length, err := conn.Read(buffer)
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Timeout() {
				fmt.Println("No more responses, client exit")
				return
			}
			panic(err)
		}
		fmt.Printf("Recieved: %s\n", string(buffer[:length]))
	}
}

// NTPの実装
func setMultiCastServer() {
	fmt.Println("Set up Multi Cast Server")
	conn, err := net.Dial("udp", "224.0.0.1:9999") // クライアント側がソケットを作成して待ち受け、そこにサーバが情報を流し込む
	if err != nil {
		panic(nil)
	}
	defer conn.Close()
	interval := 5 * time.Second
	start := time.Now()
	wait := start.Truncate(interval).Add(interval).Sub(start) // 通知時間のキリを良くするために数秒待たせる
	time.Sleep(wait)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for now := range ticker.C {
		msg := now.String()
		conn.Write([]byte(msg))
		fmt.Println("Tick: ", msg)
	}
}

func setMultiCastClient() {
	fmt.Println("Set up Multi Cast Client")
	address, err := net.ResolveUDPAddr("udp", "224.0.0.1:9999")
	if err != nil {
		panic(err)
	}
	listener, err := net.ListenMulticastUDP("udp", nil, address) // "224.0.0.1:9999" を net.UDPAddr 型に変換
	defer listener.Close()

	buffer := make([]byte, 1500)

	for {
		// length, remoteAddress, err := listener.ReadFromUDP(buffer)
		length, _, err := listener.ReadFromUDP(buffer)
		if err != nil {
			panic((err))
		}
		// fmt.Printf("Server %v\n", remoteAddress)
		fmt.Printf("Now %s\n", string(buffer[:length]))
	}
}
