package exercise

//计算每一位1出现的次数
func countDigitOne(n int) int {
	//i表示n的位数,每次进位
	//从个位开始
	//1234567为例,计算十位出现1的次数
	count := 0
	for i := 1; i <= n; i *= 10 {
		//十位的1被循环的次数
		//12345
		a := n / (i * 10)
		//十位
		//67
		b := n % (i * 10)
		//十位大于1234500,出现1的次数
		//10
		c := Min(Max(b-i+1, 0), i)
		//12345 + 10
		count += a*i + c
	}
	return count
}
func Min(args ...int) int {
	min := args[0]
	for _, item := range args {
		if item < min {
			min = item
		}
	}
	return min
}
func Max(args ...int) int {
	max := args[0]
	for _, item := range args {
		if item > max {
			max = item
		}
	}
	return max
}
