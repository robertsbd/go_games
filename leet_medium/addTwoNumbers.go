// MEDIUM

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
