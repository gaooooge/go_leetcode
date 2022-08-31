package easy

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	for i := range t {

		if t[i] == s[0] {
			s = s[1:]
		}

		if len(s) == 0 {
			return true
		}
	}

	return false

	// 作者：gaooooge
	// 链接：https://leetcode.cn/problems/is-subsequence/solution/by-gaooooge-4z83/
	// 来源：力扣（LeetCode）
	// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
}
