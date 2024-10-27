package model

import (
	"fmt"
)

type IList interface {
	Add(item Item)
	Length() int
	Display()
	Remove(item Item)
	RemovByName(name string)
	Find(item Item) bool
	FindByName(name string) Item
	Get(index int) *No
}

type No struct {
	Item Item
	Next *No
}

func NewNo() *No {
	return &No{}
}

func (n *No) Add(item Item) {
	if n == nil {
		return
	}

	if n.Item.Name == "" || n.Item.Price == 0 || n.Item.Quantity == 0 {
		n.Item = item
		return
	}

	if n.Next == nil {
		n.Next = &No{Item: item}
	} else {
		n.Next.Add(item)
	}
}

func (n *No) Length() int {
	if n == nil {
		return 0
	}
	count := 1
	current := n
	for current.Next != nil {
		count++
		current = current.Next
	}
	return count
}

func (n *No) Display() {
	current := n
	for current != nil {
		fmt.Printf("%v -> ", current.Item)
		current = current.Next
	}
}

func (n *No) Remove(item Item) {
	if n == nil || n.Next == nil {
		return
	}

	if n.Next.Item == item {
		n.Next = n.Next.Next
	} else {
		n.Next.Remove(item)
	}
}

func (n *No) RemovByName(name string) {
	if n == nil || n.Next == nil {
		return
	}

	if n.Next.Item.Name == name {
		n.Next = n.Next.Next
	} else {
		n.Next.RemovByName(name)
	}
}

func (n *No) Find(item Item) bool {
	if n == nil {
		return false
	}

	if n.Item == item {
		return true
	}

	return n.Next.Find(item)
}

func (n *No) FindByName(name string) Item {
	if n == nil {
		return Item{}
	}

	if n.Item.Name == name {
		return n.Item
	}

	return n.Next.FindByName(name)
}

func (n *No) Get(index int) *No {
	if n == nil {
		return nil
	}

	if index == 0 {
		return n
	}

	return n.Next.Get(index - 1)
}
