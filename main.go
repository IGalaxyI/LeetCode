package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow := head
	fast := head
	//Двигаем fast на n вперед
	for i := 0; i < n; i++ {
		if fast != nil {
			fast = fast.Next
		}
	}
	//Возвращяем если точка назначение это голова
	if fast == nil {
		return head.Next
	}
	//fast доходит до конца slow остается на значений удаления
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	//удаляем slow.Next
	slow.Next = slow.Next.Next

	return head
}
func removeNthFromEnd_2_0(head *ListNode, n int) *ListNode {
	L := 0
	//Вычисляем длину
	for cur := head; cur != nil; cur = cur.Next {
		L++
	}
	//Возвращяем если точка назначение это голова
	if L == n {
		return head.Next
	}

	var cur *ListNode = head
	// Проходим до момента n
	for i := 1; i < L-n; i++ {
		cur = cur.Next
	}
	//Удаляем
	cur.Next = cur.Next.Next
	return head
}
