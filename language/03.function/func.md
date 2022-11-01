# 函数

## 函数的定义

在Go中，函数的定义使用关键字`func`

- 类型相同的相邻参数可以合并;
- 可以有多个返回值，但多返回值需要加括号
- 参数是可变的

```go
func funcName(x, y int, s string, f ...float64) (int, string) {
	return x + y, s
}
```
注意：变参只能有一个，且在参数列表的末尾，其本质是传入了`slice`

## 返回值

- 命名返回参数，可由return隐式返回;
- 当命名返回参数被同名局部变量遮蔽时，需要显式返回;
- 命名返回参数允许`defer`延迟调用通过闭包读取和修改;

```go
func retTest1(x, y int) (z int) {
	z = x + y
	return //可以隐式地返回z
}

func retTest2(x, y int) (z int) {
	{
		var z = x + y //重新定义，大括号改变了z的作用域
		return z
	}
}
```

```go
func retTurn3(x, y int) (z int) {
	defer func() {
		z += 100
	}()
	
	z = x +y
	return
}

func main() {
    fmt.Println(retTurn3(1, 2))  //输出103
}
```

## 匿名函数

定义函数变量，调用变量和调用方法一样，如`funName()`
```go
fn := func() { fmt.Println("function variable") }
fn()
```

定义函数集合，以切片为例

```go
fns := []func(x int) string{
    func(x int) string { return "值为：" + string(x) },
}
fmt.Println(fns[0](100))
```

匿名函数给结构体赋值

```go
struct1 := struct {
    key string
}{
    key: "value",
}
fmt.Println(struct1.key)
```

```go
struct2 := struct {
    fn func() string
}{
    fn: func() string {
        return "hello, world"
    },
}
fmt.Println(struct2)
```

## 闭包

Go语言支持闭包，闭包是由函数及其相关引用环境组合而成的实体(即：闭包=函数+引用环境)

```go
func f(i int) func() int {
    return func() int {
        i++
        return i
    }
}
```
在函数f中返回了一个函数，返回的这个函数就是一个闭包。

这个函数中本身是没有定义变量i的，而是引用了它所在的环境（函数f）中的变量i。

## 延迟调用

关键字`defer`可以延迟调用，在return之前执行，一般用于资源释放或错误处理

```go
func deferTest() error {
	f, err := os.Create("test.txt")
	if err != nil{
		return err
	}
	
	defer f.Close()  //defer语句会在return nil之前执行
	
	f.WriteString("Hello World!")
	return nil
}
```

延迟调用函数

```go
func deferFunc() {
	x, y := 10, 20
	defer func(i int) {
		fmt.Println("defer:", i, y) //y闭包引用
	}(x)   //复制了x，x=10
	x += 10
	y += 100
	fmt.Println("x = ", x, "y = ", y)
}
```
延迟调用函数，复制了一份x的值，因此输出结果是：

```bash
x = 20 y = 120
defer: 10 120
```

## 错误处理

Go中的错误处理有独特的特点，一般使用接口`error`来处理错误，使用关键字`panic`来捕获不可修复性错误

error是一个接口类型

```go
type error interface{
	Error() string
}
```

在Go提供的内置函数中，errors.Nem()和fmt.Errorf()实现了接口，可用于创建error对象

```go
func errorTest() error {
	return errors.New("this is a error")
}
```
在包errors中，可以看到errorString实现了接口error，New()方法return具体的定义
```go
func New(text string) error {
	return &errorString{text}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
```

在Go中，内置函数panic和recover处理运行时的错误。其中，panic用于主动抛出错误，recover用来捕获panic抛出的错误。

panic有两种情况会触发

- 一种是程序出现运行时错误，直接退出，应该避免出现;
- 另一种是主动调用;

不管是显式调用panic函数，还是运行时检测到违法情况自动触发panic都会导致程序崩溃。
可以使用defer和recover捕获panic，将程序恢复正常执行。

```go
fmt.Println("this is a panic example")
defer func() {
    if r := recover(); r != nil {
        fmt.Printf("panic recover:%s", r)
    }
}()
panic("this is a panic")
```