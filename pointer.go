package main

import "fmt"

type Pointer struct {
	name string
}

func pointer() {
	var b = 5
	var a *int = &b
	*a = 90
	b = 36
	fmt.Println(*a, b)
}
