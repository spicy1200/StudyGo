package fiveWeek

import "fmt"

func checkWindow(s string, t string) string {
	//变量初始化
	lens := len(s)
	lent := len(t)
	window := make(map[byte]int, lens)
	target := make(map[byte]int, lent)
	left, right, vaild := 0, 0, 0
	res := lens
	start := -1
	for i := 0; i < lent; i++ {
		target[t[i]]++
	}
	//right右滑动
	for right < lens {
		b := s[right]
		window[b]++
		if target[b] == window[b] {
			vaild++
		}
		right++
		//left右滑动
		for vaild == len(target) {
			c := s[left]
			//是否更新结果值
			if res >= right-left {
				start = left
				res = right - left
			}
			if window[c] == target[c] {
				vaild--
			}
			window[c]--
			left++
		}
	}
	if start == -1 {
		return ""
	}
	return s[start : start+res]
}

func main()  {
	//滑动窗口
	result := checkWindow("jgh","g")
	fmt.Println(result)
}
