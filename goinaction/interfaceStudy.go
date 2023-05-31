package main

import "fmt"

//定义一个接口
type notifier interface {
	notify()
}

//定义一个类型
type user2 struct {
	name  string
	email string
}

type admin struct {
	name  string
	email string
}

// notify 是使用指针接收者实现的方法
func (u *user2) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

func (a *admin) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		a.name,
		a.email)
}

func main() {
	u := user2{"Bill", "bill@email.com"}

	lisa := admin{"Lisa", "lisa@email.com"}

	//把u换成指针类型才能编译通过
	sendNotification(&u)
	sendNotification(&lisa)
}

// sendNotification 接受一个实现了 notifier 接口的值
func sendNotification(n notifier) {
	n.notify()
}
