package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("hello world")

	i := 1

	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}
}
