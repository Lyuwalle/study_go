package main

import "fmt"

//定义接口与实现接口的简单示例

type Phone interface {
	call()
}
type NokiaPhone struct {}
type IPhone struct {}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}
func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main(){
	var phone Phone
	phone = new (NokiaPhone)
	phone.call()

	phone = new (IPhone)
	phone.call()
}
