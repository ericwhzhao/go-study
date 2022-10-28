package control

import (
	"fmt"
	"testing"
)

func TestControl(t *testing.T) {
	if num := 5; num > 0 {
		fmt.Println(num, "是一个正数")
	} else if num < 0 {
		fmt.Println(num, "是一个负数")
	}
}
