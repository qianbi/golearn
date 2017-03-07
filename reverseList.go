package main

import "flag"
import "fmt"

var (
  start int = 0
  end int = 10
)


func init() {
  flag.IntVar(&start, "start", 0, "start:int")
  flag.IntVar(&end, "end", 10, "end:int")
  flag.Parse()
}

type MyList struct {
  head *MyNode
}

type MyNode struct {
  value int
  next *MyNode
}

func createMyNode(value int, max int) *MyNode {
  if value > max {
    return nil
  }

  return &MyNode {
    value: value,
    next: createMyNode(value + 1, max),
  }
}

func reverse(node *MyNode) *MyNode {
  if node == nil || node.next == nil {
    return node
  }

  _node := reverse(node.next)
  node.next.next = node
  node.next = nil
  return _node
}
// http://www.geeksforgeeks.org/write-a-function-to-reverse-the-nodes-of-a-linked-list/
func main() {
  newList := &MyList {
    head: createMyNode(start, end),
  }
  _reversedList := &MyList {
    head: reverse(newList.head),
  }
  _newNode := _reversedList.head
  for _newNode != nil {
    fmt.Printf("%d ", _newNode.value)
    _newNode = _newNode.next
  }
}
