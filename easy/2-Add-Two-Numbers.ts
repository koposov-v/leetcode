/**
 * @link https://leetcode.com/problems/add-two-numbers/
 */
class ListNode {
  val: number
  next: ListNode | null

  constructor(val?: number, next?: ListNode | null) {
    this.val = (val === undefined ? 0 : val)
    this.next = (next === undefined ? null : next)
  }
}

const addTwoNumbers = (l1: ListNode | null, l2: ListNode | null): ListNode | null => {
  const dummyHead = new ListNode();
  let current = dummyHead;
  let carry = 0;

  while (l1 !== null || l2 !== null) {
    const x = l1 !== null ? l1.val : 0;
    const y = l2 !== null ? l2.val : 0;
    const sum = x + y + carry;

    carry = Math.floor(sum / 10);
    current.next = new ListNode(sum % 10);
    current = current.next;

    if (l1 !== null) l1 = l1.next;
    if (l2 !== null) l2 = l2.next;
  }

  if (carry > 0) {
    current.next = new ListNode(carry);
  }

  return dummyHead.next;
}


// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
const e1_l1 = arrayFillLinkedList([2, 4, 3])
const e1_l2 = arrayFillLinkedList([5, 6, 4])

// Input: l1 = [0], l2 = [0]
// Output: [0]
const e2_l1 = arrayFillLinkedList([0])
const e2_l2 = arrayFillLinkedList([0])

// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// Output: [8,9,9,9,0,0,0,1]

const e3_l1 = arrayFillLinkedList([9, 9, 9, 9, 9, 9, 9])
const e3_l2 = arrayFillLinkedList([9, 9, 9, 9])

console.log(addTwoNumbers(e1_l1, e1_l2))
// addTwoNumbers(e2_l1, e2_l1)
// addTwoNumbers(e3_l1, e3_l1)


function arrayFillLinkedList(numbers: number[], index: number = 0): ListNode | null {

  if (index === numbers.length) {
    return null;
  }

  return new ListNode(numbers[index], arrayFillLinkedList(numbers, index + 1));
}