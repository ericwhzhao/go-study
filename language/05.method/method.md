# 方法

## 方法的定义

方法与函数的区别在于，方法总是绑定对象实例，类似于面向对象语言中的成员方法；

- 隐式将实例作为第一实参，被称为接收器(receiver)
- 通过u.find()的方式调用方法
- receiver声明为非指针时，调用会产生一次拷贝；声明为指针时，指向的始终是一块内存地址
- nil是合法的接收器类型

```go
type user struct {
	name string
	age  int
	city string
}

func (u *user) find() {
    fmt.Println("user:", u.name)
}

//调用
u := user{
    //do something
}
u.find()
```

## 方法的重写

在Go中，虽然没有重写的说法，但是通过组合struct也可以实现类似重写的功能

```go
type manager struct {
	user
	title string
}

func (m *manager) find() {
    fmt.Println("manager:", m.title)
}
```
调用manager的find()方法与调用m.user的find()方法，相当于重写了父类的成员方法
```go
	m := manager{
		user{name: "hhh", age: 26, city: "beijing"},
		"manager",
	}
	m.find()
	m.user.find()
```

## 表达式

根据调用者不同，方法有两种表现形式

- method value: instance.method(args...)
- method expression: <type>.func(instance, args...)

```go
//隐式传递receiver
mValue := m.find
mValue()

//显式传递receiver
mExpression := (*manager).find
mExpression(&m) 
```
