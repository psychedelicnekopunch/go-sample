
package main


import (
	"fmt"
)


type SliceSample struct {
	List []string
}

type MapSample struct {
	List map[string]int
}


// スライスの要素は削除できないので新しく作り直す
func (s *SliceSample) Remove(index int) {
	res := []string{}
	for i, v := range s.List {
		if i == index {
			continue
		}
		res = append(res, v)
	}
	s.List = res
}


func (m *MapSample) Set(key string, value int) {
	if m.List == nil {
		m.List = map[string]int{}
	}
	m.List[key] = value
}


func (m *MapSample) Remove(key string) {
	delete(m.List, key)
}


func print(v interface{}) {
	fmt.Print(v)
	fmt.Print("\n")
}


func main() {
	// Go では配列は固定長、スライスは可変長
	// スライス (いつも使う配列)
	ss := new(SliceSample)
	print(ss.List)

	// 追加
	ss.List = append(ss.List, "test")
	print(ss.List)

	// さらに追加
	ss.List = append(ss.List, "test2")
	ss.List = append(ss.List, "test3")

	// 削除
	ss.Remove(0)
	// 存在しない要素を削除してみる
	ss.Remove(10)
	print(ss.List)


	// マップ (連想配列)
	ms := new(MapSample)
	print(ms.List)

	// 追加
	ms.Set("test", 123)
	print(ms.List)

	// さらに追加
	ms.Set("test2", 456)
	ms.Set("test3", 789)

	// 削除
	ms.Remove("test")
	// 存在しない要素を削除してみる
	ms.Remove("test4")
	print(ms.List)
}