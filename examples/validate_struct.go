package main

import (
	"fmt"
	validator "govalid"
)

type User struct {
	Name string `valid:"required"`
	Age  int    `valid:"required"`
}

func main() {
	u := User{}
	err := validator.ValidateStruct(u)
	fmt.Println(err)
}
