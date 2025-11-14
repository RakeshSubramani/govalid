package main

import (
	"fmt"
	validator "govalid"
)

func main() {
	err := validator.String("go").MinLength(3).Error()
	fmt.Println(err)
}
