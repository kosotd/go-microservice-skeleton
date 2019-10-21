package utils

import (
	"strings"
)

func Equals(str1, str2 string) bool {
	str1 = strings.ToLower(strings.Trim(str1, " "))
	str2 = strings.ToLower(strings.Trim(str2, " "))
	return strings.Compare(str1, str2) == 0
}

func Contains(str1, str2 string) bool {
	str1 = strings.ToLower(strings.Trim(str1, " "))
	str2 = strings.ToLower(strings.Trim(str2, " "))
	return strings.Contains(str1, str2)
}

func IsEmpty(str string) bool {
	return Equals("", str)
}

func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}
