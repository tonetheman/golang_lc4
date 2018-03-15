package main

import "fmt"

type m [6][6]int

func makekey(_s string) string {
	var s string
	for i := 0; i < 35; i++ {
		s += _s
	}
	return s[0:36]
}

func main() {

	var key = makekey("tony")
	fmt.Println("len of key is", len(key))
	var mat m
	fmt.Println(mat)
	for k := 0; k < 35; k++ {
		mat[k/6][k%6] = int(key[k])
	}
	fmt.Println(mat)
}
