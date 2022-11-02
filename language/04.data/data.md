# 数据存储
## 数组
- 数组的声明有三种方式;
- 数组是值传递，赋值和传参会复制整个数组;
- 数组的长度是类型的组成部分，且数组长度确定、不可变;
- 指针数组`[n]*T`和数组指针`*[n]T`的区别;

```go
//1、var声明，[]中要加数组的大小
var arr [3]int
//2、字面量声明
arr := [3]int{1, 2, 3}
//3、传入任意个数
arr := [...]int{1, 2, 3}
```
声明数组要加数组个数，或者传入`...`表示任意个数。其次，数组大小不能修改，在开发使用中不如切片（slice）方便，但数组是切片的基础。

数组的遍历，内置函数len和cap都返回数组长度
```go
for i, val := range arr {
   fmt.Println("index=", i, "value=", val)
}

for i := 0; i < len(arr); i++ {
   fmt.Println(arr[i])
}
```

## 切片
切片是数组的扩展，底层实际是数组，通过指针指向数组
```go
type slice struct {
   array unsafe.Pointer
   len   int       //元素数量
   cap   int       //最大容量
}
```
slice的创建方法有很多，以下列举几个常用的方法
```go
//1、使用var声明
var slice []int
//2、使用make关键字创建，创建[]int切片，长度len=5，容量cap=10
slice := make([]int, 5, 10)
//3、使用字面量创建
slice := []int{1,2,3,4,5}
//4、切片截取，可以从数组或者切片截取
slice := arr[2:8]
```

不管是从数组截取，还是直接创建切片，底层都会指向一个数组，对于切片的读写操作本质上都是对底层数组的操作。

切片有两个比较重要的方法，append()和copy()

```go
slice2 := append(slice1, 99)  //append()方法在slice末尾添加元素
copy(destSlice, srcSlice)     //copy()方法将一个切片复制到另一个切片
```
当slice的len小于cap时，在array[slice.high]处写数据;
当slice的cap不足以装新的元素时，进行扩容，重新分配底层数组（需要注意的是，当append多个元素时，按照全部元素数量和cap比较）;

注意：复制操作会对底层数组进行更改

## 字典
Map是常用的数据结构，在很多语言中都有实现。在Go中，Map用`哈希查找表`来实现，使用`链表`解决哈希冲突。

创建一个map，有make和字面量创建两个方法
```go
//1、使用make关键字创建
m1 := make(map[string]string)
//2、使用字面量创建map
m2 := map[string]string{}
```
常用操作：

```go
//增删改查
m["append"] = "append_data"
m["key"] = "update_data"
fmt.Println(m["append"])
delete(m, "append")
fmt.Println(m)
//获取键值对数量
fmt.Println(len(m))
```

注意；对于map，不能通过value修改value的成员变量，需要通过完整替换value或使用指针

```go
type user struct { name string}
```
```go
m1 := map[int]user {
    1: {"user1"},
}

u := m1[1]
u.name = "Tom"
m1[1] = u
fmt.Println(m1)
```

```go
m2 := map[int]*user{
    1: {
        "user2",
    },
}
m2[1].name = "Peter"
fmt.Println(m2[1])
```

## Struct

值类型，赋值和传参会复制全部内容，支持指向自身类型的指针成员

````go
f := file{
    name: "test.txt",
    size: 1024,
    attr: struct {
        perm  int
        owner int
    }{perm: 89, owner: 1},
}
fmt.Println(f.attr.perm)
````