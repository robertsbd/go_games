// Easy //

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

// longest Common Prefix

func longestCommonPrefix(strs []string) string {
    
    ch := 0;

    if len(strs) == 1 {return strs[0]}
    
    for {
        match := true;
        for i:= 0; i < len(strs) - 1; i++ {
            if strs[i] == "" {return ""}
            if len(strs[i]) - 1 < ch || len(strs[i+1]) - 1 < ch {
                match = false;
                break;
            }
            if (strs[i][ch] != strs[i+1][ch]) {
                match = false;
                break;
            }    
        }
    if match == true {ch++;} else {break};
    }
return strs[0][0:ch]
}
