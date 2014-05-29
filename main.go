package main

import _ "github.com/wfreeman/typewriters/container"

// +gen containers:"SortedSet"
type Id int32

func main() {
	/*
		ss := NewIdSortedSet(func(a, b Id) bool { return a < b })
		ss.Add(Id(1))

		fmt.Println(ss.Contains(Id(1)))
		fmt.Println(ss.Contains(Id(3)))
		ss.Add(Id(73))
		ss.Add(Id(7))
		ss.Add(Id(3))
		ss.Add(Id(2))
		ss.Add(Id(730))
		ss.Add(Id(70))
		ss.Add(Id(30))
		ss.Add(Id(20))

		for e := range ss.Iter() {
			fmt.Println(e)
		}
	*/
}
