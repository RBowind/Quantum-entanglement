# Queue

```golang
package main

import (
  "errors"
  "fmt"
)

//Queue 定义一个队列
type Queue []interface{}

//Len 队列的长度
func (queue Queue) Len() int {
  return len(queue)
}

//Push 插入队列中
func (queue *Queue) Push(value interface{}) {
  *queue = append(*queue, value)
}

//Poll 弹出
func (queue *Queue) Poll() (interface{}, error) {
  theQueue := *queue
  if theQueue.IsEmpty() {
  return nil, errors.New("queue is empty")
  }
  value := theQueue[0]
  *queue = theQueue[1:]
  return value, nil
}

//IsEmpty 判断队列是否为空
func (queue Queue) IsEmpty() bool {
  return len(queue) == 0
}

func main() {
  var myQueue Queue
  myQueue.Push(11)
  myQueue.Push("hello world")
  fmt.Println(myQueue)

  myQueue.Poll()
  fmt.Println(myQueue)

  myQueue.Poll()
  fmt.Println(myQueue.IsEmpty())
}

```
