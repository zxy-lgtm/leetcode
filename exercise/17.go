package exercise

var res []string

func letterCombinations(digits string) []string {
	res = make([]string, 0)
	if len(digits) <= 0 {
		return res
	}
	digitsMap := [10]string{
		"",     // 0
		"",     // 1
		"abc",  // 2
		"def",  // 3
		"ghi",  // 4
		"jkl",  // 5
		"mno",  // 6
		"pqrs", // 7
		"tuv",  // 8
		"wxyz", // 9
	}
	backtrack("", digits, 0, digitsMap)
	return res
}

func backtrack(tempstring, digits string, index int, digitsMap [10]string) {
	if len(tempstring) == len(digits) {
		res = append(res, tempstring)
		return
	}
	tmp := digits[index] - '0'
	ch := digitsMap[tmp]
	for i := 0; i < len(ch); i++ {
		tempstring += string(ch[i])
		backtrack(tempstring, digits, index+1, digitsMap)
		tempstring = tempstring[:len(tempstring)-1]
	}
}
