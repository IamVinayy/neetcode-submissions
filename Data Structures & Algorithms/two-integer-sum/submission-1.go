func twoSum(nums []int, target int) []int {
    a,b := 0,0
    result := [] int {}
    for i, v := range nums {
        for j := i+1; j < len(nums); j++{
            if v + nums[j] == target {
                a = i
                b = j
            }
        }
    }
    result = append(result , a)
    result = append(result, b)
    return result
}
