package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	st := make([]int, 0)
	hashmap := make(map[int]int, 0)
	n := len(nums2)
	var element int
	for i := n - 1; i >= 0; i-- {
		if len(st) == 0 {
			hashmap[nums2[i]] = -1
			st = append(st, nums2[i])
		} else {
			if st[len(st)-1] > nums2[i] {
				hashmap[nums2[i]] = st[len(st)-1]
			} else {
				element = -1
				for len(st) != 0 {
					if st[len(st)-1] <= nums2[i] {
						st = st[:len(st)-1]
					} else {
						element = st[len(st)-1]
						break
					}
				}
				hashmap[nums2[i]] = element
			}
			st = append(st, nums2[i])
		}
	}
	response := make([]int, 0)
	for _, val := range nums1 {
		num, _ := hashmap[val]
		response = append(response, num)
	}
	return response
}
