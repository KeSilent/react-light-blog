package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Case2Camel 下划线转驼峰(大驼峰)
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1) // 根据_来替换成
	name = strings.Title(name)                 // 全部大写
	return strings.Replace(name, " ", "", -1)  // 删除空格
}

// LowerCamelCase 转换为小驼峰
func LowerCamelCase(name string) string {
	name = Case2Camel(name)
	return strings.ToLower(name[:1]) + name[1:]
}

// 首字母大写
func Capitalize(s string) string {
	if s == "" {
		return ""
	}

	return strings.ToUpper(s[:1]) + s[1:]
}

func Uncaptialize(s string) string {
	if s == "" {
		return ""
	}

	return strings.ToLower(s[:1]) + s[1:]
}

func StringSliceToInt64Slice(s []string) ([]int64, error) {
	ints := make([]int64, len(s))
	for i, str := range s {
		val, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("无法将字符串 '%s' 转换成 int64: %w", str, err)
		}
		ints[i] = val
	}
	return ints, nil
}

// AdapterDiaplayToMP 将html富文本转成小程序适配
func AdapterDiaplayToMP(content string) string {
	regex := regexp.MustCompile(`<img`)
	return regex.ReplaceAllString(content, `<img width="100%"`)
}
