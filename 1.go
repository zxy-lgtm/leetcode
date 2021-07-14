func twoSum(nums []int, target int) []int {
    r := make([]int,2)
    for i := 0;i < len(nums) - 1; i++{
        for j := i+1; j < len(nums); j ++{
            if(nums[i]+nums[j] == target){
                r[0] = i
                r[1] = j
                return r
            }
        }
   }
   return r
}