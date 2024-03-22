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
    
    curNode1 := l1;
    curNode2 := l2;
    outList := List{};  
    remainder := 0;
    list1active := true;
    list2active := true;
    val1 := 0;
    val2 := 0;
    output := 0;

    for list1active == true || list2active == true {
        if list1active {
            val1 = curNode1.Val;
             if curNode1.Next == nil {list1active = false} else {curNode1 = curNode1.Next}
        } else {val1 = 0;}

        if list2active {
            val2 = curNode2.Val;
            if curNode2.Next == nil {list2active = false} else {curNode2 = curNode2.Next}
        } else {val2 = 0}

        output = remainder + val1 + val2;

        if output >= 10 {
            remainder = 1;
            output = output - 10;
        } else {remainder = 0;}

        outNode := &ListNode{
            Val: output,
        }

        if outList.head == nil {
            outList.head = outNode;
        } else {
            c := outList.head;
            for c.Next != nil {
                c = c.Next;
            }
            c.Next = outNode;
        }
    }

    if remainder == 1 {
        outNode := &ListNode{
            Val: 1,
        }

        if outList.head == nil {
            outList.head = outNode; 
        } else {
            c := outList.head;
            for c.Next != nil {
                c = c.Next;
            }
            c.Next = outNode;
        } 
    }

    return outList.head; 
}
