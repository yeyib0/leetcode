package pkg

import "strings"

func wordPattern(pattern string, s string) (b bool) {
	s2p := make(map[int32]string)
	p2s := make(map[string]int32)
	arr := strings.Split(s, " ")
	i := 0
	for _, p := range pattern {
		s := arr[i]
		if old, ok := s2p[p]; ok && old != s {
			return false
		}
		if old, ok := p2s[s]; ok && old != p {
			return false
		}
		s2p[p] = s
		i++
	}
	return true
}
