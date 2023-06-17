package logic

// leetcode 242: https://leetcode.com/problems/valid-anagram/
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sa, ta := [26]int{}, [26]int{}
	for i := 0; i < len(s); i++ {
		ac, tc := s[i]-'a', t[i]-'a'
		sa[ac] = sa[ac] + 1
		ta[tc] = ta[tc] + 1
	}

	return sa == ta
}
