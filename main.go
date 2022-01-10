package main

import "fmt"

//three part for loop. Useful for executing something a set number of times.
//three parts separated by semicolons. First part = where your index begins.
//second part = condition that allows you to continue in that loop. third =
//modifies your increment for whatever value you want.
func main() {
	for i := 10; i >= 0; i-- {
		fmt.Println("i is", i)
	}
}
