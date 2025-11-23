// カスタムエラー
// エラーは基本的にポインタで渡すようにする
// 値渡しだと比較の際に問題が生じる → 別の場所で発生（生成）したエラーでも中身が同じだと同じエラーとして扱われてしまう
package main

import "fmt"

type UserNotFound struct {
	UserName string
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found: %v", e.UserName)
}

func MyErrFunc() error {
	return &UserNotFound{UserName: "Bob"}
}

func main() {
	if err := MyErrFunc(); err != nil {
		fmt.Println(err)  // => User not found: Bob
	}
}