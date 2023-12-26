package twoway

// import "fmt"

type TwoWayLinkedCell struct {
	Value int
	Next *TwoWayLinkedCell
	Previous *TwoWayLinkedCell
}


// func (lc *TwoWayLinkedCell) AddInOrder(value int) {
// 	index := 0
// 	for lc != nil {
// 		if lc.Value > value {
// 			newTop := &TwoWayLinkedCell{Value: value}
// 			newTop.Next = lc
// 			lc = newTop
// 			fmt.Println(newTop.Next.Next.Next.Value)
			
// 		} else if lc.Value < value {
// 			if lc.Next != nil {
// 				lc = lc.Next
// 			} else {
// 				next := &TwoWayLinkedCell{Value: value}
// 				lc.Next = next
// 				return
// 			}
// 			index += 1
// 		} else {
// 			fmt.Printf("Value %d already excists at index %d\n", value, index)
// 			return
// 		}
// 	}
// }