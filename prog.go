package main

import "fmt"

type m [6][6]int

// this needs to have 36 values and
// should be a permutation
// TODO: fix this
func makekey(_s string) string {
	var s string
	for i := 0; i < 35; i++ {
		s += _s
	}
	return s[0:36]
}

func (mat *m) initState(key string) {
	for k := 0; k < 36; k++ {
		var row int = k / 6
		var col int = k % 6
		mat[row][col] = int(key[k])
	}
}

func main() {
	var key = makekey("tony")
	fmt.Println("len of key is", len(key))
	var mat m
	fmt.Println(mat)
	mat.initState(key)
	fmt.Println(mat)

	b35 := cvt_string_to_b35("a0b")
	fmt.Println(repr_b53(b35))
}
