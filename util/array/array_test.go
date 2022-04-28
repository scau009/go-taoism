package array

import (
	"fmt"
	"testing"
)

func TestDeleteItems(t *testing.T) {
	var arr1 = []string{"a", "b", "c"}
	var delArr = []string{"a"}
	result := DeleteItems(arr1, delArr)
	fmt.Println(arr1)
	fmt.Println(result)
}

func TestRemoveDuplicateItem(t *testing.T) {
	var arr1 = []string{"a", "a", "a", "a", "a"}

	result := RemoveDuplicateItem(arr1)
	fmt.Println(arr1)
	fmt.Println(result)
}
