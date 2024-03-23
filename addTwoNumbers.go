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

    func myPow(x float64, n int) float64 {

    output := 0;
    var _n float64 = 0;
    var _x float64 = 0;

    if n == 0 {return 1}

    if n < 0 { 
        _n = n * -1;
        _x = 1/x;
    } else {
        _n = n;
        _x = x;
    }

    for i := 0; i < _n - 1; i++ {
        output = output * _x;
    }
    
    return output;
}

    return outList.head; 
}

func PlainPow(x float64, n int) float64 {

    var output float64 = x;

    for i := 1; i < n; i++ {
        output = output * x;
    }

    return output;
}

func myPow(x float64, n int) float64 {

    var output float64 = 0;
    _n := 0;
    var _x float64 = 0;
    var sign float64 = 1;

    if n%2 == 1 && x < 0 {sign = -1}; 

    if n == 0 || x == 1 {return 1}

    if n == 1 {return x}

    if n < 0 { 
        _n = n * -1;
        _x = 1/x;
    } else {
        _n = n;
        _x = x;
    }

    var initial_input float64 = _x; 
    if _x < 0 {initial_input = _x*-1};

    if x == -1 {return initial_input * sign};

    // determine largest number that goes into n less than 500
    
    i := 0;

    for i = 500; _n%i != 0; i-- {}

    output = PlainPow(initial_input, i);
   
    output = PlainPow(output, int(_n / i));

    return output*sign;

}
