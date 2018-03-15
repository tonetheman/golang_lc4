package main

import (
	"strconv"
)

func repr_b53(b []byte) string {
	var ts string
	for i := 0; i < len(b); i++ {
		if b[i] >= 0 && b[i] <= 9 {
			if b[i] == 0 {
				ts += "#"
			} else if b[i] == 1 {
				ts += "_"
			} else {
				ts += strconv.Itoa(int(b[i]))
			}
		}
		if b[i] >= 10 {
			ts += strconv.Itoa(int((b[i] - 10) + 'a'))
		}
	}
	return ts
}

func cvt_string_to_b35(s string) []byte {
	var tmp []byte = make([]byte, len(s))
	for idx := range s {
		if s[idx] >= '0' && s[idx] <= '9' {
			tmp[idx] = s[idx] - 48
		}
		if s[idx] >= 'a' && s[idx] <= 'z' {
			tmp[idx] = s[idx] - 87
			// wish i had a formula for this
			// magic numbers suck
			// z = 122 -87 = 35
		}
	}
	return tmp
}
