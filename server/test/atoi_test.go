package test

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAtoi(t *testing.T) {
	id:= "1"

	fmt.Printf("%T\n",id)

	if newId,err:=strconv.Atoi(id);err==nil{
		fmt.Printf("%T\n",newId)
	}



}