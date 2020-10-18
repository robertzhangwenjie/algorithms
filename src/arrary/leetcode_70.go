package arrary

func ClimbStairs2(n int) int {
	/** 如果使用递归，则空间复杂度会太高
	动态规划思想，f(n) = f(n-1) + f(n-2)
	初始情况下f1 =1 ,f2=2 f3 = f1 + f2
	每计算一层后，将该层移掉，然后将剩下的楼层再初始化，也就是
		f2=f3, f1=f2, f3=f1 + f2
	最后返回f3
	*/

	var (
		f1 = 1
		f2 = 2
		f3 = 0
	)
	if n <= 2 {
		return n
	}

	for i := 2; i < n; i++ {
		f3 = f2 + f1
		f1 = f2
		f2 = f3
	}
	return f3
}

func ClimbStairs(n int) int {
	if n <= 3 {
		return n
	}

	return ClimbStairs(n-1) + ClimbStairs(n-2)
}
