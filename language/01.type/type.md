# 类型
## 变量

有3种方法可以声明变量。1、关键字`var`定义变量; 2、提供初始值，省略类型;3、短变量`:=`声明变量;
```go
package main

import "fmt"

func main() {
	var x int
	var y = 2.5
	z := "hello"
	x = 1
	fmt.Println(x, y, z)
}
```
在go语言中，声明了变量就必须要使用，不使用会报错。可以使用下划线`_`来代替。

## 常量

常量是编译器可确定的值，不可变，使用关键字`const`声明
```bash
const x = 1
```
注意：常量未使用不会报错

关键字`iota`可以用于定义常量组的自增常量，总是从0开始的
```go
const (
		Monday = iota //0
		Tuesday	  //1
		Wednesday //2
		Thursday  //3
		Friday    //4
		Saturday  //5
		Sunday    //6
	)
```
表达式使用`iota`有更多的用法
```go
const (
		_        = iota
		KB int64 = 1 << (10 * iota)  //iota = 1
		MB     // 1 << (10 * 2)
		GB     // 1 << (10 * 3)
		TB     // 1 << (10 * 4)
	)
	fmt.Println(KB, MB, GB, TB)
```
## 字符串

在Go中，字符串会分配到只读内存段，不能修改内容，底层实现是指向UTF-8字节数组

不能修改内容的体现如下：
```go
str := "hello"
str[0] = "H"
fmt.Println(str[0])
//报错:cannot assign to str[0] (value of type byte)
```
若想修改字符串的内容，可以整体赋值，相当于重新申请了内存，指向了新的数组
```go
str := "hello"
str = "Hello"
```
也可以转换成`[]byte`切片，脱离只读内存的限制，本质是使用了新的内存
```go
str := "hello"
bs := ([]byte)(str)
bs[0] = 'H'
fmt.Println(string(bs))
```
## 指针

支持指针类型`*T`,指针的指针`**T`,以及包含包名前缀的`*<package>.T`
- 默认值nil，没有NULL变量；
- 操作符&取变量地址，*通过指针访问目标对象；
- 直接用"."访问目标成员；

demo如下，指针p指向xy，修改p的值，xy同样改变
```go
xy := 1234
p := &xy
*p++
fmt.Println(*p, xy)
```

## 结构体

Go中的结构体类似于Java中的类，demo如下

```go
type user struct {
    name string
    age  int
    city string
}
func (u *user) hello() {
	fmt.Println("hello, my name is ", u.name)
}
```
```go
myUser := user{
    name: "peter",
    age:  18,
    city: "北京",
}
myUser.hello()
fmt.Println("name:", myUser.name, "age:", myUser.age, "city:", myUser.city)
```

