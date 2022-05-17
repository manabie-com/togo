package utils

// RemoveDuplicate remove duplicate value
func RemoveDuplicate(s []int16) []int16 {
	r := []int16{}
	m := map[int16]bool{}
	for _, v := range s {
		if !m[v] {
			r = append(r, v)
			m[v] = true
		}
	}

	return r
}
