package leetcode

type ListNode struct {
	Val int
	Next *ListNode
}
//双指针
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pa,pb:=headA,headB
	for pa!=pb{
		if pa == nil {
			pa=headB
		}else {
			pa=pa.Next
		}
		if pb==nil{
			pb=headA
		}else {
			pb=pb.Next
		}

	}
	return pa
}
//getIntersectionNode2 hashTable
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	var hashTable = map[*ListNode]bool{}
	for headA!=nil{
		hashTable[headA]=true
		headA=headA.Next
	}
	for headB!=nil{
		if hashTable[headB]{
			return headB
		}
		headB=headB.Next
	}
	return nil
}
