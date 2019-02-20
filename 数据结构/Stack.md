# Stack

```golang
package main

import (
"errors"
"fmt"
)

//Stack 定义一个栈
type Stack []interface{}

//Len 栈的高度
func (stack Stack) Len() int {
  return len(stack)
}

//Push 插入栈中
func (stack *Stack) Push(value interface{}){
  *stack = append(*stack, value)
}

//Top 栈顶
func (stack Stack) Top() (interface{}, error) {
  if len(stack) == 0 {
    return nil, errors.New("Stack is empty")
  }
  return stack[len(stack)-1], nil
}

//Pop 弹出
func (stack *Stack) Pop() (interface{}, error) {
  theStack := *stack
  if len(theStack) == 0 {
    return nil, errors.New("stack is empty")
  }
  value := theStack[len(theStack)-1]
  *stack = theStack[:len(theStack)-1]
  return value, nil
}

//IsEmpty 判断栈是否为空
func (stack Stack) IsEmpty() bool {
  return len(stack) == 0
}

func main() {
  var myStack Stack
  myStack.Push(11)
  myStack.Push("hello world")

  fmt.Println(myStack)
  myStack.Pop()
  fmt.Println(myStack.Top())

  myStack.Pop()
  fmt.Println(myStack.IsEmpty())
}

```
