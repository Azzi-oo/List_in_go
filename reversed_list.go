package main

import (
	"container/list"
)

// ReverseList - функция для реверса списка
func ReverseList(l *list.List) *list.List {
	if l == nil || l.Len() <= 1 {
		return l
	}

	reversedList := list.New()

	for element := l.Back(); element != nil; element = element.Prev() {
		reversedList.PushBack(element.Value)
	}

	return reversedList
}
