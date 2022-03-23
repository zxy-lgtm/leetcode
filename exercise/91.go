package exercise

func numDecodings(s string) int {
	f := make([]int, len(s)+1)
	f[0] = 1

	for i := 1; i <= len(s); i++ {
		if s[i-1]-'0' != 0 {
			//fmt.Println(1)
			f[i] += f[i-1]
		}

		if i > 1 && s[i-2]-'0' != 0 && ((s[i-2]-'0')*10+s[i-1]-'0') < 27 {
			f[i] += f[i-2]
		}
	}
	return f[len(s)]
}
