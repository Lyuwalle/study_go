package main

import "fmt"

// Pet 定义一个接口
type Pet interface {
	SetName(name string)	//指针方法
	Name() string			//值方法
	Category() string		//值方法
}

// Dog 只要一个数据类型的方法集合中有这 3 个方法，那么它就一定是Pet接口的实现类型。这是一种无侵入式的接口实现方式。这种方式还有一个专有名词，叫“Duck typing”
//怎样判定一个数据类型的某一个方法实现的就是某个接口类型中的某个方法呢？ 这有两个充分必要条件，一个是“两个方法的签名需要完全一致”，另一个是“两个方法的名称要一模一样”
type Dog struct {
	name string // 名字。
}

//SetName Dog类型本身的方法集合中只包含了 2 个方法，也就是所有的值方法。而它的指针类型*Dog方法集合却包含了 3 个方法（是Dog类型的方法就是*Dog类型的方法），
//所以*Dog类型就成为了Pet接口的实现类型。
func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	// 示例1。声明并初始化一个Dog类型的变量dog
	dog := Dog{"little pig"}
	_, ok := interface{}(dog).(Pet)
	fmt.Printf("Dog implements interface Pet: %v\n", ok)
	_, ok = interface{}(&dog).(Pet)
	fmt.Printf("*Dog implements interface Pet: %v\n", ok)
	fmt.Println()

	// 示例2。把dog指针类型赋值给pet
	var pet Pet = &dog
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())
}
