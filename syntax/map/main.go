package main

import (
	"fmt"
	"sort"
)

func main() {
	// mapが参照型である事の確認
	m := make(map[string]string)
	fmt.Printf("main 関数内での m 変数の値 = %d / %p\n", m, m)
	fmt.Printf("main 関数内での m 変数のアドレス = %d / %p\n", &m, &m)
	// main関数内でmは変更していないが参照渡しされるので変更される。
	addMap(m)
	fmt.Println(m)
	// copyしているので、mに影響はない
	addMapWithCopy(m)
	fmt.Println(m)

	// ---------
	// mapに対するfor
	// 学籍番号と学生名のMap
	studnetIDMap := map[int]string{
		3: "田中",
		1: "伊藤",
		2: "佐藤",
		4: "佐々木",
	}

	type student struct {
		Id int
		Name string
	}
	students := []student{}

	for k, v := range studnetIDMap {
		students = append(students, student{k, v})
		// fmt.Printfでフォーマットに従った文字列を標準出力に出せる
		fmt.Printf("Name of StudentID:%d is %s\n", k, v)
	}

	sort.Slice(students, func(i, j int) bool { return students[i].Id < students[j].Id })
	for _, e := range students {
		fmt.Printf("student Id=%d, Name=%s\n", e.Id, e.Name)
	}
}

func addMap(m map[string]string) {
	fmt.Printf("addMap 関数内での m 変数の値 = %d / %p\n", m, m)
	fmt.Printf("addMap 関数内での m 変数のアドレス = %d / %p\n", &m, &m)
	m["a"] = "あ"
}
func addMapWithCopy(m map[string]string) {
	copied := make(map[string]string)
	for k, v := range m {
		copied[k] = v
	}
	copied["i"] = "い"
}
