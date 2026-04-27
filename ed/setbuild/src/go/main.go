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

func NewSet(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}
func (v *Vector) String() string {
	str := "["
	for i := 0; i < v.size; i++ {
		if i > 0 {
			str += ", "
		}
		str += strconv.Itoa(v.data[i])
	}
	str += "]"
	return str
}
func (v *Vector) binarySearch(value int) int {
	l, r := 0, v.size-1
	for l <= r {
		m := (l + r) / 2
		if v.data[m] == value {
			return m
		} else if v.data[m] < value {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}
func (v *Vector) Insert(value int) {
	pos := v.binarySearch(value)

	if pos < v.size && v.data[pos] == value {
		return
	}

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

	for i := v.size; i > pos; i-- {
		v.data[i] = v.data[i-1]
	}

	v.data[pos] = value
	v.size++
}
func (v *Vector) Contains(value int) bool {
	for i := 0; i < v.size; i++ {
		if value == v.data[i] {
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
func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			fmt.Println(v)
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			err := v.Erase(value)
			if err != nil {
				fmt.Println(err)
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
	vec := NewSet(4)
	fmt.Println(vec.String())
}
