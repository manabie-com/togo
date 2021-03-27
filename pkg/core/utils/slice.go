package utils

func IsStringSliceContains(strSlice []string, str string) bool {
	for _, s := range strSlice {
		if s == str {
			return true
		}
	}
	return false
}

func IsInt64SliceContains(intSlice []int64, num int64) bool {
	for _, s := range intSlice {
		if s == num {
			return true
		}
	}
	return false
}

func Int64SliceUnique(intSlice []int64) []int64 {
	keys := make(map[int64]bool)
	list := []int64{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
