package main

import "fmt"

type User struct {
	Name string
}

func (u User) PrintUser() {
	fmt.Println(u.Name)
}
