package utils

import "strings"

// Is c an ASCII digit?
func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

// Is c an ASCII lower-case letter?
func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

func isASCIIUpper(c byte) bool {
	return 'A' <= c && c <= 'Z'
}

func ToASCIILower(str string) string {
	var builder = strings.Builder{}
	for i := range str {
		if isASCIIUpper(str[i]) {
			builder.WriteByte(str[i] ^ ' ')
		} else if isASCIILower(str[i]) {
			builder.WriteByte(str[i])
		}
	}
	return builder.String()
}

func SmallCamelCase(s string) string {
	if s == "" {
		return ""
	}
	t := make([]byte, 0, 32)
	i := 0
	if s[0] == '_' {
		// Need a capital letter; drop the '_'.
		t = append(t, 'X')
		i++
	}
	// Invariant: if the next letter is lower case, it must be converted
	// to upper case.
	// That is, we process a word at a time, where words are marked by _ or
	// upper case letter. Digits are treated as words.
	for ; i < len(s); i++ {
		c := s[i]
		if c == '_' && i+1 < len(s) && isASCIILower(s[i+1]) {
			continue // Skip the underscore in s.
		}
		if isASCIIDigit(c) {
			if i == 0 {
				t = append(t, '_')
			}
			t = append(t, c)
			continue
		}
		// Assume we have a letter now - if not, it's a bogus identifier.
		// The next word is a sequence of characters that must start upper case.
		if isASCIILower(c) && i != 0 {
			c ^= ' ' // Make it a capital letter.
		}
		t = append(t, c) // Guaranteed not lower case.
		// Accept lower case sequence that follows.
		for i+1 < len(s) && isASCIILower(s[i+1]) {
			i++
			t = append(t, s[i])
		}
	}
	return string(t)
}

// Lccs 最长连续公共字串
func Lccs(str1 string, str2 string) int {
	var len1, len2, ans = len(str1), len(str2), 0
	var dp = make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if str1[i] == str2[j] { // 连续
				dp[i+1][j+1] = dp[i][j] + 1
				if dp[i+1][j+1] > ans {
					ans = dp[i+1][j+1]
				}
			}
		}
	}
	return ans
}
