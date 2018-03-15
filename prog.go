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

func string_to_base35array(s string) []byte {
	var tmp []byte = make([]byte, len(s))
	for idx := range s {
		if s[idx] == "0" {
			tmp[idx] = 0
		}
		if s[idx] == "1" {
			tmp[idx] = 1
		}
	}
	return tmp
}

func main() {
	var key = makekey("tony")
	fmt.Println("len of key is", len(key))
	var mat m
	fmt.Println(mat)
	mat.initState(key)
	fmt.Println(mat)
}
