#### What is a Linked List?

A linked list is a linear data structure where elements are stored in nodes. Each node contains two parts:

- Data: The value stored in the node.
- Next: A pointer/reference to the next node in the sequence.

#### Why Use a Linked List?

Linked lists are useful when you need a dynamic data structure that can grow and shrink in size. They are particularly useful for:

- Efficient insertions and deletions.
- Implementing other data structures like stacks and queues.

#### Structure of the Linked List

##### Node

A Node represents an element in the linked list.

```go
type Node struct {
    data int
    next *Node
}
```

- `data`: Stores the value of the node.
- `next`: A pointer to the next node in the list. If it's the last node, next will be nil.


##### LinkedList

A LinkedList manages the nodes.

```go
type LinkedList struct {
    head *Node
}
```

- `head`: A pointer to the first node in the list. If the list is empty, head will be nil.

#### Operations on a Linked List

##### Insert at the Back

Adds a new node with the given data at the end of the list.

```go
func (list *LinkedList) insertAtBack(data int) {
    newNode := &Node{data: data, next: nil}

    if list.head == nil {
        list.head = newNode
        return
    }

    current := list.head
    for current.next != nil {
        current = current.next
    }

    current.next = newNode
}
```

Explanation: A new node is created with the given data. If the list is empty, the new node becomes the head. Otherwise, the new node is added at the end of the list.

##### Insert at the Front

Adds a new node with the given data at the beginning of the list.

```go
func (list *LinkedList) insertAtFront(data int) {
	if list.head == nil {
		newNode := &Node{data: data, next: nil}
		list.head = newNode
		return
	}

	newNode := &Node{data: data, next: list.head}
	list.head = newNode
}
```

Explanation: A new node is created with the given data. If the list is empty, the new node becomes the head. Otherwise, the new node is added at the front of the list.