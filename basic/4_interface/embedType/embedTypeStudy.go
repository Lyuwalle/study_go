package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

// User defines a User in the program.
type User struct {
	name  string
	email string
}

// notify implements a method that can be called via a value of type User.
// *User类型就是notifier接口的实现类型
func (u *User) notify() {
	fmt.Printf("Sending User email to %s<%s>\n",
		u.name,
		u.email)
}

// admin represents an admin User with privileges.
type admin struct {
	User  // Embedded Type 嵌入类型
	level string
}

// main is the entry point for the application.
func main() {
	// Create an admin User.
	ad := admin{
		User: User{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// We can access the inner type's method directly.
	ad.User.notify()

	// The inner type's method is promoted. 调用嵌入类型的方法
	ad.notify()

	sendNotification(&ad)
}

func sendNotification(n notifier) {
	n.notify()
}
