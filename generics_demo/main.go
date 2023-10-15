//package main
//
//import "fmt"
//
//type MySlice[T int | float32] []T
//
//func (s MySlice[T]) Sum() T {
//	var sum T
//	for _, value := range s {
//		sum += value
//	}
//	return sum
//}
//
//type Int interface {
//	int | int8 | int16 | int32 | int64
//}
//type Uint interface {
//	uint | uint8 | uint16 | uint32
//}
//type Float interface {
//	float32 | float64
//}
//
//func Add[T Int | Uint | Float](a T, b T) T {
//	return a + b
//}
//
//type Queue[T interface{}] struct {
//	elements []T
//}
//
//func (q *Queue[T]) Put(value T) {
//	q.elements = append(q.elements, value)
//}
//func main() {
//	var s1 MySlice[int] = []int{1, 2, 3, 4}
//	fmt.Println(s1.Sum())
//
//	var s2 MySlice[float32] = []float32{1.0, 2.0, 3.0, 4.0}
//	fmt.Println(s2.Sum())
//
//	fmt.Println(Add(4, 5))
//	fmt.Println(Add(4.0, 5.0))
//
//	var q1 Queue[string]
//	q1.Put("hello")
//
//}

package main

import (
	"errors"
	"fmt"
)

func Add[T int | float64](a, b T) T {
	return a + b
}

type MySlice[T int | float64] []T

func (s MySlice[T]) Sum() T {
	var sum T
	for _, value := range s {
		sum += value
	}
	return sum
}

type Queue[T int | float64] struct {
	elements []T
}

func (q *Queue[T]) Put(value T) {
	q.elements = append(q.elements, value)
}
func (q *Queue[T]) Pop() (value T, err error) {
	if len(q.elements) < 1 {
		return value, errors.New("no value")
	}
	value = q.elements[0]
	q.elements = q.elements[1:]
	return value, nil
}
func main() {

	fmt.Println(Add(1, 2))
	fmt.Println(Add(1.0, 2.0))

	var s MySlice[int] = []int{1, 2, 3}
	fmt.Println(s.Sum())

	var q1 Queue[int]
	q1.Put(1)
	q1.Put(2)
	v, err := q1.Pop()
	fmt.Println(v, err)
}
