package split

import "strings"

func Split(s, sep string) (result []string) {
	for {
		i := strings.Index(s, sep)
		if i == -1 {
			result = append(result, s)
			break
		}
		result = append(result, s[:i])		//string切片还是string
		s = s[i+len(sep):]
	}
	return
}
