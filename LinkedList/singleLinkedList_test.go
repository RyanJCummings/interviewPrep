package main

import (
    "testing"
    "reflect"
)

func TestCreateLinkedList(t *testing.T) {
    list := createLinkedList(1)
    want := 1
    if want != list.head.key {
        t.Errorf("Got %v, Want %v\n", list.head.key, want)
    }
}

func TestCreateNode(t *testing.T) {
    node := createNode(5)
    want := 5
    if node.key != want {
        t.Errorf("Got %v, Want %v\n", node.key, want)
    }
}

func TestInsertNode(t *testing.T) {
    l := createLinkedList(2)

    // At beginning of list
    l.insertNode(createNode(1), l.head)

    // At end of list
    l.insertNode(createNode(4), l.head)

    // In the middle of the list
    l.insertNode(createNode(3), l.head)

    want := []int{1, 2, 3, 4}
    got := l.visitNodes(l.head, []int{})

    if !reflect.DeepEqual(want, got) {
        t.Errorf("Got: %v, Want: %v\n", got, want)
    }
}

// deleteNode_Test
func TestDeleteNode(t *testing.T) {
    l := createLinkedList(1)
    l.insertNode(createNode(2), l.head)
    l.insertNode(createNode(3), l.head)
    l.insertNode(createNode(4), l.head)
    l.insertNode(createNode(5), l.head)

    // At beginning of list
    l.deleteNode(1, l.head)
    got := l.visitNodes(l.head, []int{})
    want := []int{2, 3, 4, 5}
    if !reflect.DeepEqual(got, want) {
        t.Errorf("Got: %v, Want: %v\n", got, want)
    }

    // At end of list
    l.deleteNode(5, l.head)
    got = l.visitNodes(l.head, []int{})
    want = []int{2, 3, 4}
    if !reflect.DeepEqual(got, want) {
        t.Errorf("Got: %v, Want: %v\n", got, want)
    }

    // In the middle of the list
    l.deleteNode(3, l.head)
    got = l.visitNodes(l.head, []int{})
    want = []int{2, 4}
    if !reflect.DeepEqual(got, want) {
        t.Errorf("Got: %v, Want: %v\n", got, want)
    }
}
