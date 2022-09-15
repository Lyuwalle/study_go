package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {

	bob := user{"Bob", "bob@gmail.com"}
	bob.notify()

	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	bob.changeEmail("bob@hotmail.com")
	bob.notify()

	lisa.changeEmail("lisa@hotmail.com")
	lisa.notify()
}
