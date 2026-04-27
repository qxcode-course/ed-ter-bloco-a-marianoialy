package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewMultiSet(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}
func (v *Vector) String() string {
	return "[" + Join(v.data[0:v.size], ", ") + "]"
}
func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}
func (v *Vector) Insert(value int) {
	if v.size == v.capacity {
		newCapacity := 1
		if v.capacity > 0 {
			newCapacity = v.capacity * 2
		}
		newData := make([]int, newCapacity)
		copy(newData, v.data[:v.size])
		v.data = newData
		v.capacity = newCapacity
	}
	found, pos := v.Search(value)
	if found {
		pos++
	}
	for i := v.size; i > pos; i-- {
		v.data[i] = v.data[i-1]
	}
	v.data[pos] = value
	v.size++
}
func (v *Vector) Search(value int) (bool, int) {
	left := 0
	right := v.size - 1
	result := v.size

	for left <= right {
		mid := (left + right) / 2

		if v.data[mid] == value {
			result = mid
			left = mid + 1
		} else if v.data[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
			result = mid
		}
	}
	if result < v.size && v.data[result] == value {
		for result+1 < v.size && v.data[result+1] == value {
			result++
		}
		return true, result
	}
	return false, left
}
func (v *Vector) Contains(value int) bool {
	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			return true
		}
	}
	return false
}
func (v *Vector) Erase(value int) error {
	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			for j := i; j < v.size-1; j++ {
				v.data[j] = v.data[j+1]
			}
			v.size--
			return nil
		}
	}
	return fmt.Errorf("value not found")
}
func (v *Vector) Count(value int) int {
	count := 0
	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			count++
		}
	}
	return count
}
func (v *Vector) Unique() int {
	if v.size == 0 {
		return 0
	}
	count := 1
	for i := 1; i < v.size; i++ {
		if v.data[i] != v.data[i-1] {
			count++
		}
	}
	return count
}
func (v *Vector) Clear() {
	v.size = 0
}
func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			// value, _ := strconv.Atoi(args[1])
			// ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms)
		case "erase":
			value, _ := strconv.Atoi(args[1])
			err := ms.Erase(value)
			if err != nil {
				fmt.Println(err)
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Contains(value))
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
	vec := NewMultiSet(4)
	fmt.Println(vec.String())
}
