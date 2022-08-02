package test

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	var a *int
	a = new(int)
	*a = 10
	fmt.Println(a)
}

func TestPoint2(t *testing.T) {
	var a int
	var a2 = &a
	*a2 = 1
	fmt.Println(a)
}

type Cat struct {
}

func (*Cat) miao() {
	fmt.Println("miao")
}

func TestStruct(t *testing.T) {
	cat := new(Cat)
	(*cat).miao()
}
