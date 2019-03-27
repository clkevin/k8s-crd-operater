package v1

import "fmt"

type Get struct {
	Base
}


func (Get) Operate()string{
	fmt.Println("get")

	return "Get";
}