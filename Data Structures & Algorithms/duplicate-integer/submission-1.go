func hasDuplicate(nums []int) bool {
      var result bool = false
    for i := 0; i< len (nums); i++ {
    	for j := i+1; j<len(nums); j++ {
    		if nums[i] == nums[j]{
              result = true
    			return result 
    		}
    	}
    }
  return result
}
