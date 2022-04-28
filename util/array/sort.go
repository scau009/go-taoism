package array

import (
	"fmt"
	"strconv"
)

/**
选择排序,map[任意值]数字（浮点型或整型）
返回：map的key数组
*/
func SelectedSort(mappings map[interface{}]interface{}) ([]interface{}, error) {
	var arr = []interface{}{}
	for mapping, _ := range mappings {
		arr = append(arr, mapping)
	}

	mappingsLen := len(mappings)

	for i := 0; i < mappingsLen; i++ {
		compareScore, err := strconv.ParseFloat(fmt.Sprintf("%v", mappings[arr[i]]), 64)
		if err != nil {
			return nil, err
		}
		idx := i
		for j := i + 1; j < mappingsLen; j++ {
			currentScore, err := strconv.ParseFloat(fmt.Sprintf("%v", mappings[arr[j]]), 64)
			if err != nil {
				return nil, err
			}
			if currentScore > compareScore {
				idx = j
				compareScore = currentScore
			}
		}
		arr[i], arr[idx] = arr[idx], arr[i]
	}

	return arr, nil
}

func SelectedSortInterfaceIntMap(mappings map[interface{}]int) ([]interface{}, error) {
	var arr = []interface{}{}
	for mapping, _ := range mappings {
		arr = append(arr, mapping)
	}

	mappingsLen := len(mappings)

	for i := 0; i < mappingsLen; i++ {
		compareScore, err := strconv.ParseFloat(fmt.Sprintf("%v", mappings[arr[i]]), 64)
		if err != nil {
			return nil, err
		}
		idx := i
		for j := i + 1; j < mappingsLen; j++ {
			currentScore, err := strconv.ParseFloat(fmt.Sprintf("%v", mappings[arr[j]]), 64)
			if err != nil {
				return nil, err
			}
			if currentScore > compareScore {
				idx = j
				compareScore = currentScore
			}
		}
		arr[i], arr[idx] = arr[idx], arr[i]
	}

	return arr, nil
}

func SelectedSortStringIntMap(mappings map[string]int) []string {
	var arr = []string{}
	for mapping, _ := range mappings {
		arr = append(arr, mapping)
	}

	mappingsLen := len(mappings)

	for i := 0; i < mappingsLen; i++ {
		score := mappings[arr[i]]
		idx := i
		for j := i + 1; j < mappingsLen; j++ {
			if mappings[arr[j]] > score {
				idx = j
				score = mappings[arr[j]]
			}
		}
		arr[i], arr[idx] = arr[idx], arr[i]
	}

	return arr
}
