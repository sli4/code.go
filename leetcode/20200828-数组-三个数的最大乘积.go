package leetcode
// https://leetcode-cn.com/problems/maximum-product-of-three-numbers/
// 三种可能：三大正、二负一正、三大负
func maximumProduct(nums []int) int {
	a,b,c := 0,0,0
	e,f := 0,0
	flag := false
	g,h,k := 0,0,0
	for i:= 0; i<len(nums); i++ {
		if nums[i] < 0 {
			if nums[i] < e {
				f = e
				e = nums[i]
			} else if nums[i]<f {
				f = nums[i]
			}

			if flag==false {
				// 最大负数
				if g == 0 {
					g = nums[i]
					continue
				}
				if h==0 {
					if nums[i] > g {
						h = g
						g =nums[i]
					} else {
						h = nums[i]
					}
					continue
				}
				if k==0 {
					if nums[i] > h {
						k=h
						if nums[i] > g {
							h = g
							g = nums[i]
						} else {
							h = nums[i]
						}
					} else {
						k = nums[i]
					}
					continue
				}

				if nums[i]> g {
					k=h
					h=g
					g=nums[i]
				} else if nums[i] > h {
					k=h
					h=nums[i]
				} else if nums[i] >k {
					k = nums[i]
				}
			}
		} else {
			flag = true
			if nums[i] > a {
				c = b
				b = a
				a = nums[i]
			} else if nums[i] > b {
				c = b
				b = nums[i]
			} else if nums[i] > c {
				c = nums[i]
			}
		}
	}
	if flag {
		if a*e*f > a*b*c {
			return a*e*f
		}
		return a*b*c
	} else {
		return g*h*k
	}

}

func maxNegative(a,b,c,v int) {
	if a==0 {
		a = v
		return
	}
}