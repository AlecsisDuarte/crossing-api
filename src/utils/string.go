package utils

import "strconv"

func IsEmpty(s *string) (b bool) {
	return len(*s) == 0
}

func IsNotEmpty(s *string) (b bool) {
	return !IsEmpty(s)
}

func IsInt(s *string) (b bool) {
	if IsEmpty(s) {
		return false
	}
	if _, err := strconv.Atoi(*s); err != nil {
		return false
	}
	return true
}

func IsNotInt(s *string) (b bool) {
	return !IsInt(s)
}

func ToInt(s *string) (v int, err error) {
	return strconv.Atoi(*s)
}

func ToBool(s *string) (b bool, err error) {
	return strconv.ParseBool(*s)
}
