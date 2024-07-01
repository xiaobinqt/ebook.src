class Solution:
    def reverseBetween(self, head: Optional[ListNode], left: int, right: int) -> Optional[ListNode]:
        if head is None:
            return head

        dummy = ListNode()
        dummy.next = head
        prev = dummy

        for _ in range(left-1):
            prev = prev.next

        curr = prev.next
        for _ in range(right-left):
            next_node = curr.next
            curr.next = curr.next.next
            next_node.next = prev.next
            prev.next = next_node

        return dummy.next
