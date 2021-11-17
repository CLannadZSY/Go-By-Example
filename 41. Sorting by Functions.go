package main

import (
	"fmt"
	"sort"
)

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"c", "bb", "aaa"}

	// 按照长度排序
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)

	// 按照首字母大小
	sort.Strings(fruits)
	fmt.Println(fruits)
}
