package data

import (
	"fmt"
	"testing"
)

type user struct {
	name string
}

func TestArray(t *testing.T) {
	//1、var声明，[]中要加数组的大小
	var arr1 [3]int
	//2、字面量声明
	arr2 := [3]int{1, 2, 3}
	//3、传入任意个数
	arr3 := [...]int{1, 2, 3}

	fmt.Println(arr1, arr2, arr3)
}

func TestSlice(t *testing.T) {
	//1、使用var声明
	//var slice []int
	//2、使用make关键字创建，创建[]int切片，长度len=5，容量cap=10
	//slice := make([]int, 5, 10)
	//3、使用字面量创建
	slice := []int{1, 2, 3, 4, 5}
	//4、切片截取，可以从数组或者切片截取
	//slice := arr[2:8]

	s2 := append(slice, 999)
	fmt.Println(s2)

	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := data[8:]
	s1 := data[:5]
	fmt.Println(s)
	fmt.Println(s1)
	copy(s1, s)
	fmt.Println(s1)
	fmt.Println(data)
}

func TestMap(t *testing.T) {
	//m1 := make(map[string]string, 10)
	m := map[string]string{
		"key": "value",
	}
	//增删改查
	m["append"] = "append_data"
	m["key"] = "update_data"
	fmt.Println(m["append"])
	delete(m, "append")
	fmt.Println(m)
	//获取键值对数量
	fmt.Println(len(m))

	m1 := map[int]user{
		1: {"user1"},
	}

	u := m1[1]
	u.name = "Tom"
	m1[1] = u
	fmt.Println(m1)

	m2 := map[int]*user{
		1: {
			"user2",
		},
	}
	m2[1].name = "Peter"
	fmt.Println(m2[1])
}

type file struct {
	name string
	size int
	attr struct {
		perm  int
		owner int
	}
}

func TestStruct(t *testing.T) {
	f := file{
		name: "test.txt",
		size: 1024,
		attr: struct {
			perm  int
			owner int
		}{perm: 89, owner: 1},
	}
	fmt.Println(f.attr.perm)
}
