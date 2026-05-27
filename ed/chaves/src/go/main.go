package main

import "fmt"

func main() {
	queue := NewQueue[string]()
	for letra := 'A'; letra <= 'P'; letra++ {
		queue.Enqueue(string(letra))
	}
	for queue.items.Len() > 1 {
		time1 := queue.Dequeue()
		time2 := queue.Dequeue()
		var gols1 int
		var gols2 int
		fmt.Scan(&gols1, &gols2)
		if gols1 > gols2 {
			queue.Enqueue(time1)
		} else {
			queue.Enqueue(time2)
		}
	}
	fmt.Println(queue.Dequeue())
}
