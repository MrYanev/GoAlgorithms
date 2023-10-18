package module01

import (
	"math"
)

func BaseToDec(value string, base int) int {
	var ans int
	for idx, val := range value {
		res := math.Pow(float64(base), float64((len(value) - idx)))
		ans += int(val) * int(res)
	}
	return ans
}
