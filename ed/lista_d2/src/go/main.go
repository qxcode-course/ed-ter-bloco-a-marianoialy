package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	info int
	next *Node
	prev *Node
	root *Node
}
type Dlist struct {
	head *Node
	size int
}

func NewLList() *Dlist {
	dlist := &Dlist{}
	dlist.head = &Node{}
	dlist.head.next = dlist.head
	dlist.head.prev = dlist.head
	return dlist
}
func PushBack(list *Dlist, value int) {
	insert(list, list.head, value)
}
func PushFront(list *Dlist, value int) {
	insert(list, list.head.next, value)
}
func String(list *Dlist) {
	if list.size == 0 {
		fmt.Println("[]")
		return
	}

	fmt.Print("[")
	for it := list.head.next; it != list.head; it = it.next {
		fmt.Printf("%d", it.info)
		if it.next != list.head {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}
func insert(list *Dlist, node *Node, value int) {
	newNode := &Node{
		info: value,
		root: list.head,
	}

	prev := node.prev

	newNode.prev = prev
	newNode.next = node

	prev.next = newNode
	node.prev = newNode

	list.size++
}
func (l *Dlist) Front() *Node {
	if l.size == 0 {
		return nil
	}
	return l.head.next
}
func (l *Dlist) Back() *Node {
	if l.size == 0 {
		return nil
	}
	return l.head.prev
}
func (n *Node) Next() *Node {
	if n.next == n.root {
		return nil
	}
	return n.next
}

func (n *Node) Prev() *Node {
	if n.prev == n.root {
		return nil
	}
	return n.prev
}
func Clear(list *Dlist) {
	list.head.next = list.head
	list.head.prev = list.head
	list.size = 0
}
func (l *Dlist) Search(value int) *Node {
	for it := l.head.next; it != l.head; it = it.next {
		if it.info == value {
			return it
		}
	}
	return nil
}
func Remove(list *Dlist, node *Node) *Node {
	if node == nil || node == list.head {
		return nil
	}

	next := node.next

	node.prev.next = node.next
	node.next.prev = node.prev

	list.size--

	if next == list.head {
		return nil
	}
	return next
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			String(ll)
		case "size":
			// fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				PushBack(ll, num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				PushFront(ll, num)
			}
		case "pop_back":
			// ll.PopBack()
		case "pop_front":
			// ll.PopFront()
		case "clear":
			Clear(ll)
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.info)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.info)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.info = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				insert(ll, node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				Remove(ll, node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
