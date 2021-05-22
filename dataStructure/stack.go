package dataStructure

import "fmt"

func CreateStack() []int {
	stack := make([]int, 3)
	return stack
}

func PushStack(stack []int, item int) {
	stack = append(stack, item)
	fmt.Printf("Pushed item ->%v to %v", item, stack)
}

func PopStack(stack []int) {

}
