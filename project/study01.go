package main

import "fmt"

func main() {
	// 最初が小文字ならprivate,大文字ならpublic
	// msg := "hello"

	/**
	* 変数には宣言した後に初期値が入る
	* int 0
	* string ""
	*/

	const (
		sun = iota // const識別子
		mon // 1
		tue // 2
	)

	a := 5
	var pa *int //int型のポインタ
	pa = &a // &a = aのアドレス
	// paの領域にあるデータの値= *pa
	fmt.Println(pa)
	fmt.Println(*pa)

	fmt.Println(getMessage("まい"))

	// 変数に関数を入れる
	f := func(name1, name2 string)(string, string) {
		return name2, name1
	}
	fmt.Println(f("mai", "teru"))

	// 即時関数的な
	func(name1, name2 string) {
		fmt.Println(name1 + name2)
	}("mai", "teru")

	// 配列の初期化
	a := [...]{1, 2, 3}
}

// 変数をreturnできる
func getMessage(name string) (msg string) {
	msg = name + "さん"
	return
}