package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	data *list.List
}

func Push(elem interface{}, queue *list.List) {
	queue.PushBack(elem)
}

func Pop(queue *list.List) interface{} {
	if queue.Len() == 0 {
		return nil
	}
	frontElement := queue.Front()
	value := frontElement.Value
	queue.Remove(frontElement)
	return value
}

func printQueue(queue *list.List) {
	for element := queue.Front(); element != nil; element = element.Next() {
		fmt.Print(element.Value)
	}
}
