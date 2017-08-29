package helper

func In_slice_p(sl []interface{}, p interface{}) int {
	if p == nil || len(sl) == 0 {
		return -1
	}

	for i := 0; i < len(sl); i++ {
		if sl[i] == p {
			return i + 1
		}
	}
	return -1
}
