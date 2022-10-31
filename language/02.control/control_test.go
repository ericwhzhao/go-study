package control

import (
	"fmt"
	"testing"
)

func TestIf(t *testing.T) {
	if num := 5; num > 0 {
		fmt.Println(num, "是一个正数")
	} else if num < 0 {
		fmt.Println(num, "是一个负数")
	}
}

func TestFor(t *testing.T) {
	str := "hello, go-study"
	n := len(str)
	//1、普通用法
	for i := 0; i < n; i++ {
		fmt.Println(string(str[i]))
	}

	//2、类似于while语句
	//相当于其他语言的while (n > 0) {}
	for n > 0 {
		//do something
		n--
	}

	//3、while(true)的替代
	for {
		//do something
		fmt.Println("hello, go-study")
	}
}

func TestRange(t *testing.T) {
	//定义切片mySlice
	mySlice := []int{1, 2, 3}
	//range的基本使用
	for i, val := range mySlice {
		fmt.Println("i:", i, "val:", val)
	}
	//可忽略不使用的返回值，使用下划线
	for i, _ := range mySlice {
		fmt.Println("index:", i)
	}
	//也可以直接忽略第二个返回值，但是只支持string、array、slice、map
	for i := range mySlice {
		fmt.Println("index:", i)
	}

	a := [3]int{0, 1, 2}
	for i, v := range a {
		if i == 0 {
			a[1], a[2] = 999, 999
		}
		a[i] = v + 100
		fmt.Println(a)
		a[i] = v + 200
		fmt.Println(a)
	}
	fmt.Println(a)
}

func TestSwitch(t *testing.T) {
	i := 1
	//1、基本使用
	switch i {
	case 1:
		fmt.Println("a")
		fallthrough
	case 0:
		fmt.Println("b")
	default:
		fmt.Println("default")
	}
	//2、当作if语句使用，省略条件表达式
	switch {
	case i > 0:
		fmt.Println(i, "是一个正数")
	case i < 0:
		fmt.Println(i, "是一个负数")
	default:
		fmt.Println("零值")
	}
	//3、可以在是switch中初始化，一般结合判断使用
	switch j := true; {
	case j == true:
		fmt.Println("true")
	default:
		fmt.Println("false")
	}
}

func TestGoto(t *testing.T) {
	var i int = 10
LOOP:
	for i < 20 {
		if i == 15 {
			i = i + 1
			goto LOOP
		}
		fmt.Printf("i的值为：%d\n", i)
		i++
	}
}
