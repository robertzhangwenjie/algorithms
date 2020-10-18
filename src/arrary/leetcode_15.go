package arrary

import "sort"

func ThreeSum1(nums []int) [][]int {
	/**
	解法 a+b+c=0 ,先排序，然后使用双指针左右 往中间逼近
	1. 对数组进行排序，a<b<c，则a<0,c>0
	2. 遍历数组，假设3个指针k,i,j分别代表a、b、c，寻找每一个k的组合解，如果k与k-1相同，则跳过k
		i和j位于数组的头和尾，如果 nums[k] + nums[i] + nums[j] <0，则nums[i]太小，i++ ,此时如果nums[i-1]== nums[i],则跳过
						如果 nums[k] + nums[i] + nums[j] >0，则nums[j]太大，j--,此时如果nums[j] == nums[j+1]，则跳过
	*/
	var results [][]int

	sort.Ints(nums)
	for k := 0; k < len(nums); k++ {
		if nums[k] > 0 {
			break
		}
		if k >= 1 && nums[k] == nums[k-1] {
			continue
		}
		for i, j := k+1, len(nums)-1; i < j; {
			res := nums[i] + nums[j] + nums[k]
			if res > 0 {
				j--
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			} else if res < 0 {
				i++
				for i < j && nums[i] == nums[i-1] {
					i++
				}
			} else {
				results = append(results, []int{nums[k], nums[i], nums[j]})
				i++
				j--
				for i < j && nums[i] == nums[i-1] {
					i++
				}
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			}
		}
	}
	return results
}

func ThreeSum(nums []int) [][]int {
	var results [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					res := []int{nums[i], nums[j], nums[k]}
					results = append(results, res)
				}
			}
		}
	}
	return results
}
