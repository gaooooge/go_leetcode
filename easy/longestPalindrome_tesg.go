package easy

func longestPalindrome(s string) string {
	maxStr := ""

	//奇数
	for i := range s {
		index1 := i
		index2 := i
		for index1 >= 0 && index2 <= len(s)-1 && s[index1] == s[index2] {
			if len(s[index1:index2+1]) > len(maxStr) {
				maxStr = s[index1 : index2+1]
			}
			index1--
			index2++
		}
	}

	//偶数
	for i := range s {
		index1 := i
		index2 := i + 1
		for index1 >= 0 && index2 <= len(s)-1 && s[index1] == s[index2] {
			if len(s[index1:index2+1]) > len(maxStr) {
				maxStr = s[index1 : index2+1]
			}
			index1--
			index2++
		}
	}

	return maxStr
}
