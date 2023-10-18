package module01

import "fmt"

func DecToBase(dec, base int) string {
	var ans string
	for dec != 0 {
		tmp := dec % base
		ans = fmt.Sprintf("%X", tmp) + ans
		dec = dec / base
	}
	return ans
}
