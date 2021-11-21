package leetcode

func findComplement(num int) int {
	var HighBit int
	for i := 30; i > 0; i-- {
		if num>>i == 1 {
			HighBit = i
			break
		}
	}
	flag := 1<<(HighBit+1) - 1
	return flag ^ num
}
