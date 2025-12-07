package main

import (
	"fmt"
	"time"
)

func main() {
	S1()
	S2()
	S3()
	S4()
}

// goroutineでの並列実行
func S1() {
	fmt.Println("S1() is running")
	go S1Sub()
	time.Sleep(2 * time.Second) // S1がS1Subよりも先に終了してしまう（S1Subが終了するのを待たないため）
	fmt.Println("S1() is finished")
}

func S1Sub() {
	fmt.Println("S1Sub() is running")
	time.Sleep(time.Second)
	fmt.Println("S1Sub() is finished")
}

// チャネルの基礎
func S2() {
	tasks := make(chan string, 10)
	tasks <- "cmake .."
	tasks <- "cmake . --build Debug"
	task, ok := <-tasks
	if ok {
		fmt.Printf("task=%#v\n", task)
	}
}

// 終了まで待機
func S3() {
	fmt.Println("S3 is running")
	done := make(chan bool)
	go func() {
		fmt.Println("S3 is finished")
		done <- true
	}()
	<-done
	fmt.Println("All tasks are finished")
}

// ポーリング
func S4() {
	fmt.Println("S4 is running")
	tasks := make(chan string, 10)
	go func() {
		time.Sleep(time.Second)
		tasks <- "completed"
	}()
	for { // ポーリング
		select {
		case task := <-tasks:
			fmt.Printf("task=%#v\n", task)
			fmt.Println("S4 is finished")
			return
		default:
			break
		}
	}
}
