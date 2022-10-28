package main

import "fmt"

func main() {
	//1、变量
	var x int
	var y = 2.5
	z := "hello"
	x = 1

	const s = 9
	fmt.Println(x, y, z)
	fmt.Println(y)
	fmt.Println(z)

	//2、常量
	const (
		Monday    = iota //0
		Tuesday          //1
		Wednesday        //2
		Thursday         //3
		Friday           //4
		Saturday         //5
		Sunday           //6
	)

	const (
		_        = iota
		KB int64 = 1 << (10 * iota)
		MB
		GB
		TB
	)
	fmt.Println(1 << (20))

	//3、字符串
	str := "hello"
	//str = "Hello"
	bs := ([]byte)(str)
	bs[0] = 'H'
	fmt.Println(string(bs))

	xy := 1234
	p := &xy
	*p++
	fmt.Println(*p)
	fmt.Println(xy)

	//4、结构体

	myUser := user{
		name: "peter",
		age:  18,
		city: "北京",
	}
	myUser.hello()
	fmt.Println("name:", myUser.name, "age:", myUser.age, "city:", myUser.city)
}

type user struct {
	name string
	age  int
	city string
}

func (u *user) hello() {
	fmt.Println("hello, my name is ", u.name)
}
