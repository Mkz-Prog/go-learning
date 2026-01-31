package main

import "fmt"

type User struct {
	age int
	name string
}

var users = []User{{200, "Yoda"}, {18, "Michael"}, {20, "John"}}

func FindOldest(users []User) User {
	var oldest = users[0]
	for _, val := range(users) {
		if val.age > oldest.age { oldest = val }
	}
	return oldest
}

func HappyBithday(u *User) {
	u.age ++
}

func main() {
	HappyBithday(&users[0])
	var oldest User = FindOldest(users)
	fmt.Printf("Oldest person is %s, he is %d years old.\n", oldest.name, oldest.age)
}