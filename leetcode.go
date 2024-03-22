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

func isPalindrome(x int) bool {

    var str = strconv.Itoa(x)
    var start = 0;
    var end = len(str) - 1;

    for i:= 0; i < end; i++ {
        if start > end{
            return false;
        }
        if start == end{
            return true;
        }
        if str[start] != str[end]{
            return false;
        }
        start++;
        end--;
    }
    return true;
}

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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type List struct {
    head *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    
    num1:= 0;
    num2:= 0;
    i := 0;
    curNode := l1;

    for {
        num1 += curNode.Val*int(math.Pow10(i));
        if curNode.Next == nil {
            break;
        } else {
            curNode = curNode.Next; 
            i++;
        }
    }

    fmt.Println(num1)

    curNode = l2;
    i = 0;

    for {
        num2 += curNode.Val*int(math.Pow10(i));
        if curNode.Next == nil {
            break;
        } else {
            curNode = curNode.Next; 
            i++;
        }
    }

    fmt.Println(num2)

    outNum := strconv.Itoa(num1 + num2);

    outList := List{};

    for i := len(outNum) - 1; i >= 0; i-- {
       node := &ListNode{
           Val: int(outNum[i] - '0'),
       } 
       if outList.head == nil {
            outList.head = node; 
        } else {
            c := outList.head;
            for c.Next != nil {
                c = c.Next;
            }
            c.Next = node;
        }
   }

    return outList.head; 
}
