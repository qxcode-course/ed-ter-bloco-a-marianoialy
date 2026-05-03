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
}
type Dlist struct {
	head *Node
	size int
}

func NewDlist() *Dlist {
	dlist := &Dlist{}
	dlist.head = &Node{}
	dlist.head.next = dlist.head
	dlist.head.prev = dlist.head
	return dlist
}
func insert(list *Dlist, A *Node, value int) {
	B := A.prev

	C := &Node{
		info: value,
		next: A,
		prev: B,
	}

	B.next = C
	A.prev = C
	list.size++
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
func Size(list *Dlist) int {
	return list.size
}
func Clear(list *Dlist) {
	list.head.next = list.head
	list.head.prev = list.head
	list.size = 0
}
func PopFront(list *Dlist) {
	if list.size == 0 {
		return
	}

	first := list.head.next
	next := first.next

	list.head.next = next
	next.prev = list.head

	list.size--
}
func PopBack(list *Dlist) {
	if list.size == 0 {
		return
	}

	last := list.head.prev
	before := last.prev

	before.next = list.head
	list.head.prev = before

	list.size--
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewDlist()

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
			fmt.Println(Size(ll))
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
			PopBack(ll)
		case "pop_front":
			PopFront(ll)
		case "clear":
			Clear(ll)
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
