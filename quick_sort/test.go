package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	left := 0
	res := 0
	//for i, ch := range s {
	for i := 0; i < len(s); i++ {
		//m[ch] = 1
		//fmt.Println(i, s[i])
		//fmt.Printf("%T\n", s[i])
		ch := s[i]
		if elem, ok := m[ch]; ok {
			if m[ch] > left {
				left = m[ch] + 1
			}
		}
		if res < i-left+1 {
			res = i - left + 1
		}
		m[ch] = i

	}
	return res
}
*/
/*
func main() {
	test_str := "hello&world&12&12313"
	//lengthOfLongestSubstring(test_str)
	test_split := strings.Split(test_str, "&")
	fmt.Println(test_split)
	for _, v := range test_split {
		fmt.Println(strings.Index(v, "1"))
		test_num, err := strconv.Atoi(v)
		fmt.Println(test_num, err)
	}
}
*/
/*
func validIPAddress(queryIP string) string {
	if index := strings.Index(queryIP, "."); index != -1 {
		ip_slice := strings.Split(queryIP, ".")
		if len(ip_slice) != 4 {
			return "Neither1"
		}
		for _, str := range ip_slice {
			if strings.Index(str, "0") == 0 && len(str) != 1 {
				return "Neither2"
			}
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println(err, str, num)
				return "Neither3"
			}
			if num > 255 || num < 0 {
				return "Neither4"
			}
		}
		return "IPv4"
	} else {
		ip_slice := strings.Split(queryIP, ":")
		if len(ip_slice) != 8 {
			return "Neither"
		}
		for _, str := range ip_slice {
			if len(str) > 4 || len(str) < 1 {
				return "Neither"
			}
			_, err := strconv.ParseUint(str, 16, 16)
			if err != nil {
				return "Neither"
			}
		}
		return "IPv6"
	}
	return "Neither"
}
*/

func validIPAddress(queryIP string) string {
	if index := strings.Index(queryIP, "."); index != -1 {
		slice_str := strings.Split(queryIP, ".")
		if len(slice_str) != 4 {
			fmt.Println(slice_str)
			return "Neither"
		}
		for _, str := range slice_str {
			if strings.Index(str, "0") != -1 && len(str) != 1 {
				fmt.Println("zero", str)
				return "Neither"
			}
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println(err, num, str)
				return "Neither"
			}
			fmt.Println(num, str, err)
			if num > 255 || num < 0 {
				fmt.Println(num)
				return "Neither"
			}
		}
		return "IPv4"
	} else {
		slice_str := strings.Split(queryIP, ":")
		if len(slice_str) != 8 {
			return "Neither"
		}
		for _, str := range slice_str {
			if len(str) > 4 {
				return "Neither"
			}
			_, err := strconv.ParseUint(str, 16, 16)
			if err != nil {
				return "Neither"
			}

		}
		return "IPv6"
	}
	return "Neither"
}

/*
func main() {
	// queryIP := "172.16.254.1"
	//queryIP := "2001:0db8:85a3:0:0:8A2E:0370:7334"
	// queryIP := "20EE:FGb8:85a3:0:0:8A2E:0370:7334"
	// queryIP := "192.0.0.1"
	queryIP := "1.1.1."
	str_res := validIPAddress(queryIP)
	fmt.Println(str_res)

	// str_test := "1.1.1."
	// slice_test := strings.Split(str_test, ".")
	// fmt.Println(slice_test, len(slice_test))
	// test1, err := strconv.Atoi(slice_test[3])
	// if err != nil {
	// 	fmt.Println(test1, err)
	// }
	// fmt.Println(test1, err)

}
*/
/*
type ListNode struct {
	Val  int
	Next *ListNode
}
func isPalindrome(head *ListNode) bool {
	slice_num := make([]int, 0)
	for head != nil {
		append(slice_num, head.Val)
	}
	left := 0
	right := len(slice_num) - 1
	for left < right {
		if slice_num[left] != slice_num[right] {
			return false
		}
		left++
		right--
	}
	return true
}
*/
func main() {
	slice_num := make([]int, 0)
	//var slice_num []int
	append(slice_num, 0)
	fmt.Println(slice_num)
}
