package method

import (
	"fmt"
	"testing"
)

type user struct {
	name string
	age  int
	city string
}

func (u *user) find() {
	fmt.Println("user:", u.name)
}

type manager struct {
	user
	title string
}

func (m *manager) find() {
	fmt.Println("manager:", m.title)
}

func TestMethod1(t *testing.T) {
	u := user{
		name: "xxx",
		age:  19,
		city: "shanghai",
	}
	u.find()
}

func TestMethod2(t *testing.T) {
	m := manager{
		user{name: "hhh", age: 26, city: "beijing"},
		"manager",
	}
	m.find()
	m.user.find()
}

func TestMethod3(t *testing.T) {
	//表达式
	//定义一个实例
	m := manager{
		user{name: "hhh", age: 26, city: "beijing"},
		"管理员",
	}
	mValue := m.find
	mValue() //隐式传递receiver

	mExpression := (*manager).find
	mExpression(&m) //显式传递receiver
}
