package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	x, s := funcName(1, 2, "func", 2.2, 2.2, 2.2, 2.2)
	fmt.Println(x, s)

	fmt.Println(retTurn3(1, 2))

	fn := func() { fmt.Println("function variable") }
	fn()

	fns := []func(x int) string{
		func(x int) string { return "值为：" + string(x) },
	}
	fmt.Println(fns[0](100))

	d := struct {
		key string
	}{
		key: "value",
	}
	fmt.Println(d.key)

	struct2 := struct {
		fn func() string
	}{
		fn: func() string {
			return "hello, world"
		},
	}
	fmt.Println(struct2)

	fmt.Println(errorTest())

	panicTest()

}

func funcName(x, y int, s string, f ...float64) (int, string) {
	return x + y, s
}

func retTest1(x, y int) (z int) {
	z = x + y
	return
}

func retTest2(x, y int) (z int) {
	{
		var z = x + y
		return z
	}
}

func retTurn3(x, y int) (z int) {
	defer func() {
		z += 100
	}()

	z = x + y
	return
}

func deferTest() error {
	f, err := os.Create("test.txt")
	if err != nil {
		return err
	}

	defer f.Close()

	f.WriteString("Hello World!")
	return nil
}

func deferFunc() {
	x, y := 10, 20
	defer func(i int) {
		fmt.Println("defer:", i, y)
	}(x)
	x += 10
	y += 100
	fmt.Println("x=", x, "y=", y)
}

func errorTest() error {
	return errors.New("this is a error")
}

func panicTest() {
	fmt.Println("this is a panic example")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic recover:%s", r)
		}
	}()
	panic("this is a panic")
}
