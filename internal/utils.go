package internal

func intersection(a, b []int) []int {
	counter := make(map[int]struct{})
	var result []int

	for _, v := range a {
		counter[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := counter[v]; ok {
			result = append(result, v)
		}
	}
	return result
}
