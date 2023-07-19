package main

import (
	"errors"
	"fmt"
)

//声明一个函数类型, 参数和结果都是int类型
type operate func(x int, y int) int

type calculateFunc func(x int, y int) (int, error)

//高阶函数（满足下面两个条件之一的函数：接受其他的函数作为参数传入，或者把其他的函数作为结果返回）
func calculate(x int, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}

//高阶函数，把其他的函数作为结果返回
//定义一个匿名的、calculateFunc类型的函数并把它作为结果值返回。
func genCalculator(op operate) calculateFunc {
	return func(x int, y int) (int, error) {
		if op == nil { return 0, errors.New("invalid operation") }
		return op(x, y), nil
	}
}

func main() {
	plus := func(x, y int) int {
		return x + y
	}
	//调用高阶函数
	fmt.Println(calculate(1, 2, plus))
	fmt.Println(genCalculator(plus)(1, 2))
}
