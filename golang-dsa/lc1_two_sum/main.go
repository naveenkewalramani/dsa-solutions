package main
import("fmt")
// https://leetcode.com/problems/two-sum/description/

func twoSum(nums []int, target int)[]int {
	var response []int
	hashmap := make(map[int]int)
	for index, value := range nums {
		hashmap[value] = index
	}
	for index, value := range nums{
		item, ok := hashmap[target-value]
		if !ok {
			continue
		}

		if item == value{
			continue
		}else{
			response = append(response, []int{index, item}...)
			break
		}
	}
	return response
}

func main(){
	nums := []int{2,7,11,15}
	target := 9
	fmt.Println(twoSum(nums, target))
}