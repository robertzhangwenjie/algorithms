package arrary

func CompareSlice(s1 []int, s2 []int) bool {
	/**
	比较两个切片中的元素是否相同
	*/
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		isExist := false
		len_s2 := len(s2)
		for j := 0; j < len_s2; j++ {
			if s1[i] == s2[j] {
				isExist = true
				s2 = append(s2[:j], s2[j+1:]...)
				break
			}

		}
		if !isExist {
			return false
		}
	}
	return true
}
