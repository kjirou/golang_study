package main

import (
	"fmt"
)

func main() {
	n := 100
	fmt.Printf("main 関数内での n 変数の値 = %d / %p\n", n, n)
	fmt.Printf("main 関数内での n 変数のアドレス = %d / %p\n", &n, &n)
	// 値渡し
	// コピーされるので、元のnに変化はない
	returnValue := increment(n)
	fmt.Printf("Value of n is %d\n", n)
	fmt.Printf("Return Value of increment is %d\n", returnValue)
	// 参照渡し
	// 戻り値を受け取らなくても渡した変数が書き換わる
	incrementWithPointer(&n)
	fmt.Printf("Value of n is %d\n", n)
}

func increment(n int) int {
	fmt.Printf("increment 関数内での n 変数の値 = %d / %p\n", n, n)
	fmt.Printf("increment 関数内での n 変数のアドレス = %d / %p\n", &n, &n)
	return n + 1
}
func incrementWithPointer(n *int) {
	fmt.Printf("incrementWithPointer 関数内での n 変数の値 = %d / %p\n", n, n)
	fmt.Printf("incrementWithPointer 関数内での n 変数のアドレス = %d / %p\n", &n, &n)
	*n++
}
