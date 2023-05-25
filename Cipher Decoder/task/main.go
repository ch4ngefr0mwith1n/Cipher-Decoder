package main

import "fmt"

func main() {
	var g, p int
	_, err := fmt.Scanf("g is %d and p is %d", &g, &p)
	if err != nil {
		panic(err)
	}

	fmt.Printf("g=%d and p=%d", g, p)
}
