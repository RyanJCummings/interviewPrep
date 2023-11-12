// TODO:
    // Error handling
    // Interfaces/Generics
package main

import (
    "fmt"
)

type Node struct {
    key     int
    next    *Node
}

type LinkedList struct {
    head *Node
}

func createNode(key int) *Node {
    return &Node {
        key:    key,
        next:   nil,
    }
}

func createLinkedList(key int) *LinkedList {
    return &LinkedList{createNode(key)}
}

func (l *LinkedList) insertNode(newNode, compareNode *Node) {
    // base case: duplicates (none allowed for simplicity)
    if newNode.key == compareNode.key {
        fmt.Printf("%+v\n is already a member of the list", newNode.key)
    // beginning of list
    } else if newNode.key < compareNode.key {
        newNode.next = compareNode
        l.head = newNode
    } else if newNode.key > compareNode.key {
        // end of list
        if compareNode.next == nil {
            compareNode.next = newNode
        // new node belongs between two nodes
        } else if newNode.key < compareNode.next.key {
            newNode.next = compareNode.next
            compareNode.next = newNode
        // recursively call this function and increment compareNode by 1
        } else {
            l.insertNode(newNode, compareNode.next)
        }
    }
}

func (l *LinkedList) visitNodes(node *Node, nodes []int) []int {
    if node.next != nil {
        nodes = append(nodes, node.key)
        return l.visitNodes(node.next, nodes)
    }
    nodes = append(nodes, node.key)
    return nodes
}

func (l *LinkedList) deleteNode (target int, node *Node) {
    // if target node is head
    if target == l.head.key {
        l.head = l.head.next
    // if target is not in the list
    } else if node.next == nil {
        fmt.Printf("%+v\n was not found in the list", node.key)
    // if the next node is the target
    } else if node.next.key == target {
        node.next = node.next.next
    // if the next node is not the target
    } else if node.next.key < target {
        l.deleteNode(target, node.next)
    }
}

func (l *LinkedList) printList(node *Node) {
    // print the entire list
    l.printNode(node)
    if node.next != nil {
        l.printList(node.next)
    }
}

func (l *LinkedList) printNode(node *Node) {
    // print a node
    fmt.Println(node.key)
}

// Just for fun as a review of pointers
func (l *LinkedList) getListNodeValues_Pointer(node *Node, nodes *[]int) {
    *nodes = append(*nodes, node.key)
    if node.next != nil {
        l.getListNodeValues_Pointer(node.next, nodes)
    }
}

