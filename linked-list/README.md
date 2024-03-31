# Linked List

A linked list is a linear data structure where elements are stored in nodes. Each node contains two parts: data and a reference (or pointer) to the next node in the sequence. Unlike arrays, linked lists do not have a fixed size in memory and can dynamically grow or shrink as needed.

## Operations

### Insertion
Inserting an element into a linked list involves creating a new node and updating the references appropriately to include the new node in the list. There are several common insertion scenarios:
- Insert at the beginning of the list
- Insert at the end of the list
- Insert after a specific node

### Deletion
Deleting an element from a linked list involves finding the node containing the element to be deleted and updating the references to exclude that node from the list. There are also several common deletion scenarios:
- Delete from the beginning of the list
- Delete from the end of the list
- Delete a specific node by value

### Traversal
Traversing a linked list means iterating through the nodes to access or manipulate each element. This operation is often used to display the elements of the list or perform some operation on each element.

### Searching
Searching for an element in a linked list involves traversing the nodes to find a specific element by its value.

![linked list](https://static.wikia.nocookie.net/leetcode/images/9/96/Linked_list.png/)

## Pseudo code for linked list basic operations

#### Insertion
```text
InsertAtBeginning(data):
    newNode = CreateNode(data)
    newNode.next = head
    head = newNode
```

```text
InsertAtEnd(data):
    newNode = CreateNode(data)
    if head is NULL:
        head = newNode
    else:
        current = head
        while current.next is not NULL:
            current = current.next
        current.next = newNode
```

```text
InsertAfterNode(prevNode, data):
    if prevNode is NULL:
        return
    newNode = CreateNode(data)
    newNode.next = prevNode.next
    prevNode.next = newNode
```
### Deletion

```text
DeleteFromBeginning():
    if head is NULL:
        return
    head = head.next
```

```text
DeleteFromEnd():
    if head is NULL:
        return
    if head.next is NULL:
        head = NULL
        return
    current = head
    while current.next.next is not NULL:
        current = current.next
    current.next = NULL
```

```text
DeleteNode(data):
    if head is NULL:
        return
    if head.data == data:
        head = head.next
        return
    current = head
    while current.next is not NULL:
        if current.next.data == data:
            current.next = current.next.next
            return
        current = current.next
```

### Traversal

```text
TraverseList():
    current = head
    while current is not NULL:
        print(current.data)
        current = current.next
```
### Searching
```text
Search(data):
    current = head
    while current is not NULL:
        if current.data == data:
            return true
        current = current.next
    return false
```