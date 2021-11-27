package offer

// 在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

// 示例 1:

// 输入：s = "abaccdeff"
// 输出：'b'
// 示例 2:

// 输入：s = ""
// 输出：' '
//

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func FirstUniqChar(s string) byte {
	size := len(s)
	if size == 0 {
		return ' '
	}
	// map[rune]{first show index}
	exceptSet := make(map[rune]struct{})
	m := make(map[rune]int)
	for k, v := range s {
		if _, ok := exceptSet[v]; ok {
			continue
		}
		if _, ok := m[v]; ok {
			delete(m, v)
			exceptSet[v] = struct{}{}
			continue
		}
		m[v] = k
	}
	first := size + 1
	if len(m) != 0 {
		for _, v := range m {
			if v < first {
				first = v
			}
		}
	}
	if first > size {
		return ' '
	}
	return s[first]
}
