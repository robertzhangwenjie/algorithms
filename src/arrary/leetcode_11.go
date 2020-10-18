package arrary

func MaxArea2(height []int) int {
	var maxArea = 0
	var minHeight = 0
	for i, j := 0, len(height)-1; i < j; {
		if height[i] < height[j] {
			minHeight = height[i]
			i++
		} else {
			minHeight = height[j]
			j--
		}
		area := minHeight * (j - i + 1)
		if maxArea < area {
			maxArea = area
		}
	}
	return maxArea

}

func MaxArea(height []int) int {
	/*
		解题思路1：对所有的线进行两两组合，算出他们的面积，然后取出最大值
		1. 维护两个指针，枚举所有的组合 ，取出最大的一个
	*/
	var (
		area    = 0
		maxArea = 0
	)
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			if height[i] > height[j] {
				area = (j - i) * height[j]
			} else {
				area = (j - i) * height[i]
			}
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}
