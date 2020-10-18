package arrary

//leetcode 283题，移动零

func MoveZeroes2(nums []int) {
	// 解题方法 ，
	// 1. 使用两个指针，i和j。i为快指针，j为慢指针。
	//		j为移动0后的数组指针，i为当前数组的指针
	// 2. 遍历切片，如果遇到非0值，就将索引i的值赋予给索引为j的元素，并将j指针向前移动一位
	//	如果此时i和j的值相同，那就不需要移动，如果i和j的值不同，那么移动i后需要将i指向的元素的值置为0

	var j = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			if i != j {
				nums[i] = 0
			}
			j++
		}
	}

}

func MoveZeroes(nums *[]int) {
	//解题方法
	//1. 计算0的个数
	//2. 将非0值往前移动，移动的次数取决于前面有多少个0
	//3. 最后，根据0的个数，将切片的最后的n位数赋值0

	var counter = 0
	for i, num := range *nums {
		if num != 0 {
			(*nums)[i-counter] = num
		} else {
			counter++
		}
	}

	for i := 0; i < counter; i++ {
		(*nums)[len(*nums)-i-1] = 0
	}
}
