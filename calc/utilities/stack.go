package utilities

type Stack struct {
	data []interface{}
}

func NewStack() *Stack {
	return &Stack{}
}

func (stack *Stack) Push(item interface{}) {
	stack.data = append(stack.data, item)
}

func (stack *Stack) Peek() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	return stack.data[len(stack.data)-1]
}

func (stack *Stack) IsEmpty() bool {
	if len(stack.data) == 0 {
		return true
	}
	return false
}

func (stack *Stack) Pop() {
	if stack.IsEmpty() {
		return
	}
	stack.data = stack.data[:len(stack.data)-1]
}
