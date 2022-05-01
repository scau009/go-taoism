package string_tool

import (
	"fmt"
	"testing"
)

// 调试时需要计算md5
func TestMd5(t *testing.T) {
	result := Md5("testOrder0010.0nullTony YanMwKKEago2iUJAAzv6rcDCRRadu7hnE9hflbuF4KL")
	fmt.Println(result)
}

//
func TestCrc32(t *testing.T) {
	// Amazon Redshift
	// 625007 8P307 00020
	result := Crc32("XX")
	fmt.Println(result) //3230741573
	result = Crc32("ML")
	fmt.Println(result) //3230741573
}

func TestF(t *testing.T) {
	input := "2fafj9f4-wfjaivn7&0vs"

	fmt.Println(GetNumFromString(input))
}
