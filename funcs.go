// find two elements in an array that add up to a target number

func twoSum(nums []int, target int) []int {

    var m = make(map[int]int)
    var target2 = 0

    for i := 0; i < len(nums); i++ {
        m[nums[i]] = i;
    }

    for j := 0; j < len(nums); j++ {
        target2 = target - nums[j];
        indx, ok := m[target2];
        if ok && j != indx {
            return []int{j, indx}
        }
    }
    return []int{-99,-99}
 }   
