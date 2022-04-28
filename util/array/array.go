package array

import (
	"reflect"
	"strconv"
	"strings"
)

func Int2Str(items []int) []string {
	strIds := make([]string, 0)
	for _, item := range items {
		strIds = append(strIds, strconv.Itoa(item))
	}
	return strIds
}

/**
删除数组中重复的内容
*/
func RemoveDuplicateItem(items []string) []string {
	result := make([]string, 0, len(items))
	temp := map[string]bool{}
	for _, item := range items {
		if _, ok := temp[item]; !ok {
			temp[item] = true
			result = append(result, item)
		}
	}
	return result
}

/**
根据内容，删除数组中的内容
*/
func DeleteItems(originalItems []string, deleteItems []string) []string {
	items := []string{}
	for _, originalItem := range originalItems {
		if !Contains(deleteItems, originalItem) {
			items = append(items, originalItem)
		}
	}

	return items
}

/**
根据数组下标，删除数组内容，适用于同一时间要删除多个下标的值
*/
func DeleteItemsByIndex(originalItems []string, indexs []int) []string {
	items := []string{}
	for index, originalItem := range originalItems {
		if !ContainsInt(indexs, index) {
			items = append(items, originalItem)
		}
	}

	return items
}

/**
根据下标删除数组的元素
*/
func DeleteStringItemsByIdx(items []string, idx int) []string {
	if idx < 0 || len(items) < idx || len(items) == 0 {
		return items
	}

	if idx > 0 && idx < len(items) {
		return append(items[:idx], items[idx+1:]...)
	}

	if idx == 0 && len(items) == 1 {
		return []string{}
	} else if idx == 0 {
		return items[idx+1:]
	}

	return items[:idx]
}

/**
判断数组中是否包含指定的内容
*/
func Contains(items []string, str string) bool {

	for _, item := range items {
		if item == str {
			return true
		}
	}

	return false
}

// HasItemIn 判断数组items是否有元素在targetItems 里面
func HasItemIn(items []string, targetItems []string) bool {

	for _, item := range items {
		if Contains(targetItems, item) {
			return true
		}
	}

	return false
}

/**
搜索传入数组的下标值
@example
	Search([]string{"one", "two"}, "two") = 0
	Search([]string{"one", "two"}, "three") = nil
	Search(map[string]int{"one":1, "two":2}, 1) = "one"
	Search(map[string]int{"one":1, "two":2}, 3) = nil

*/
func Search(items interface{}, searchVal interface{}) interface{} {
	if items == nil {
		return nil
	}

	itemTypel := reflect.TypeOf(items)
	itemsVal := reflect.ValueOf(items)

	switch itemTypel.Kind() {
	case reflect.Slice, reflect.Array:

		for i := 0; i < itemsVal.Len(); i++ {
			if itemsVal.Index(i).Interface() == searchVal {
				return i
			}
		}
	case reflect.Map:
		for _, item := range itemsVal.MapKeys() {
			if itemsVal.MapIndex(item).Interface() == searchVal {
				return item.Interface()
			}
		}
	}

	return nil
}

func ContainsInt(items []int, num int) bool {
	for _, item := range items {
		if item == num {
			return true
		}
	}

	return false
}

/**
过滤字符串中空的数组
@example
	array.Filter([]string{"abc"," cc",""})
	得到：
	[]string{"abc"," cc"}
*/
func Filter(originalItems []string, filters ...string) []string {
	if len(filters) == 0 {
		filters = []string{""}
	}
	return DeleteItems(originalItems, filters)
}

/**
遍历字符串数组，对数组每个元素使用同样的方法得到新的字符串数组
@example
	array.Map([]string{"abc"," cc"," bb "}, strings.TrimSpace)
	得到：
	[]string{"abc","cc","bb"}
*/
func Map(originalItems []string, callback func(text string) string) []string {
	items := []string{}
	for _, originalItem := range originalItems {
		items = append(items, callback(originalItem))
	}

	return items
}

func ToLower(originalItems []string) []string {
	items := []string{}
	for _, originalItem := range originalItems {
		items = append(items, strings.ToLower(originalItem))
	}

	return items
}

/**
将字符串数组按长度分割成二维数组
*/
func Chunk(items []string, chunkSize int) (chunks [][]string) {
	for chunkSize < len(items) && chunkSize > 0 {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}

	return append(chunks, items)
}

/**
整数数组最大值
*/
func Max(items []int) int {
	var m int
	for i, e := range items {
		if i == 0 || e > m {
			m = e
		}
	}
	return m
}

/**
整数数组最小值
*/
func Min(items []int) int {
	var m int
	for i, e := range items {
		if i == 0 || e < m {
			m = e
		}
	}
	return m
}

/**
切割 int数组 to 最大长度为 size 的二维数组
*/
func SliceBySize(slice []int, size int) [][]int {
	result := make([][]int, 0)
	subSlice := make([]int, 0)
	for _, value := range slice {
		subSlice = append(subSlice, value)
		if len(subSlice) >= size {
			result = append(result, subSlice)
			subSlice = make([]int, 0)
		}
	}
	if len(subSlice) > 0 {
		result = append(result, subSlice)
	}
	return result
}

// DiffStringSlice 对比数组 arr1与arr2，返回arr1中比arr2多的元素
func DiffStringSlice(arr1 []string, arr2 []string) []string {
	if len(arr2) == 0 {
		return arr1
	}
	// 转换成map
	tmpMap := map[string]bool{}
	for _, val := range arr2 {
		tmpMap[val] = true
	}
	diffArr := []string{}
	for _, val := range arr1 {
		if _, exist := tmpMap[val]; !exist {
			diffArr = append(diffArr, val)
		}
	}
	return diffArr
}
