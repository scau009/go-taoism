package string_tool

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/syyongx/php2go"
	"hash/crc32"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

/*
将string中的数字转化为float
*/
func StringToNum(num interface{}) float64 {
	var result float64 = 0
	switch num.(type) {
	case string:
		result, _ = strconv.ParseFloat(num.(string), 64)
		break
	case float64:
		result = num.(float64)
		break
	case float32:
		result = float64(num.(float32))
		break
	case int:
		tmpNum := num.(int)
		result = float64(tmpNum)
		break
	case int64:
		tmpNum := num.(int64)
		result = float64(tmpNum)
	}
	return result
}

/*
判断字符串是否包含中文字符
*/
func IsHasChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

/**
使用crc32 加密字符串
*/
func Crc32(str string) string {
	crc32q := crc32.MakeTable(crc32.IEEE)
	uint32Result := crc32.Checksum([]byte(str), crc32q)
	return fmt.Sprint(uint32Result)
}

/**
不区分大小写替换字符串
*/
func ReplaceCaseInsensitive(search, replace, subject string) string {
	re := regexp.MustCompile(`(?i)` + search)
	return re.ReplaceAllString(subject, replace)
}

func StripTags(content string) string {
	re := regexp.MustCompile(`<(.|\n)*?>`)
	return re.ReplaceAllString(content, "")
}

func MergeSpaces(content string) string {
	re := regexp.MustCompile("\\s{2,}")

	return re.ReplaceAllString(content, "\n")

}

func GetRandomStr(length int) string {
	var str = "abcdefghijklmnopqrstuvwxyz"
	var result = ""
	for i := 0; i < length; i++ {
		str = php2go.StrShuffle(str)
		result = result + str[0:1]
	}

	return result
}

func GetRandomStrWithNumber(length int) string {
	var str = "abcdefghijklmnopqrstuvwxyz0123456789"
	var result = ""
	for i := 0; i < length; i++ {
		str = php2go.StrShuffle(str)
		result = result + str[0:1]
	}

	return result
}

func GetRandomNumber(length int) string {
	var str = "0123456789"
	var result = ""
	for i := 0; i < length; i++ {
		str = php2go.StrShuffle(str)
		result = result + str[0:1]
	}

	return result
}

func Md5(str string) string {
	data := md5.Sum([]byte(str))
	return hex.EncodeToString(data[:])
}

func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str) //

	isFirst := false
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			isFirst = true
		}

		if isFirst {
			if vv[i] >= 97 && vv[i] <= 122 { // 97 a , 122 z
				vv[i] -= 32 // a - 32 = A
			}
			isFirst = false
		} else {
			if vv[i] >= 65 && vv[i] <= 90 { // 65 A , 90 Z
				vv[i] += 32 // A + 32 = a
			}
		}

		upperStr += string(vv[i])

		if vv[i] == 32 { // 32 为 空格
			isFirst = true
		}
	}
	return upperStr
}

func GetNumFromString(str string) string {
	valid := regexp.MustCompile("[0-9]")
	result := valid.FindAllStringSubmatch(str, -1)
	stringArr := make([]string, 0)
	for _, r := range result {
		stringArr = append(stringArr, r[0])
	}
	return strings.Join(stringArr, "")
}

func Int2Str(items []int) []string {
	strIds := make([]string, 0)
	for _, item := range items {
		strIds = append(strIds, strconv.Itoa(item))
	}
	return strIds
}
