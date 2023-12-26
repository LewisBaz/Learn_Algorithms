package oneway

import (
	"fmt"
)

type OneWayLinkedCell struct {
	Value int
	Next *OneWayLinkedCell
}

func (lc *OneWayLinkedCell) AddInEnd(value int) {
	for lc != nil {
		if lc.Next != nil {
			lc = lc.Next
		} else {
			next := &OneWayLinkedCell{Value: value}
			lc.Next = next
			return
		}
	}
}

func (lc *OneWayLinkedCell) AddAfter(target int, value int) {
	isFound := false
	for lc != nil {
		if lc.Value != target {
			lc = lc.Next
		} else if lc.Value == target {
			isFound = true
			next := &OneWayLinkedCell{Value: value}
			next.Next = lc.Next
			lc.Next = next
			return
		}
	}
	if !isFound {
		fmt.Printf("Target %d not found\n", target)
	}
}

func (lc *OneWayLinkedCell) FindNextAfter(value int) *OneWayLinkedCell {
	for lc != nil {
		if lc.Value == value {
			return lc.Next
		} else {
			lc = lc.Next
		}
	}
	return nil
}

func (lc *OneWayLinkedCell) Remove(value int) *OneWayLinkedCell {
	for lc != nil && lc.Value == value {
		newHead := lc.Next
		return newHead
	}
	for lc != nil && lc.Next != nil {
		if lc.Next.Value == value {
			lc.Next = lc.Next.Next
			return lc
		} else {
			lc = lc.Next 
		}
	}
	return nil
}

func (lc *OneWayLinkedCell) ShowAll() {
	index := 0
	for lc != nil {
		fmt.Printf("Value at index %d is %d\n", index, lc.Value)
		lc = lc.Next
		index += 1
	}
}