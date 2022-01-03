package exercise

import "strconv"

func addStrings(num1 string, num2 string) (ans string) {
	tag := 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || tag != 0; i, j = i-1, j-1 {
		x, y := 0, 0
		if i >= 0 {
			x = int(num1[i] - '0')
		}

		if j >= 0 {
			y = int(num2[j] - '0')
		}

		res := x + y + tag
		ans = strconv.Itoa(res%10) + ans
		tag = res / 10
	}
	return
}
