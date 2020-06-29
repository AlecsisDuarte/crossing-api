package utils

import "strconv"

// IsEmpty validates if the string is empty
func IsEmpty(s *string) (b bool) {
	return len(*s) == 0
}

// IsNotEmpty validates if the string is not empty
func IsNotEmpty(s *string) (b bool) {
	return !IsEmpty(s)
}

// IsInt validates if the string is a number
func IsInt(s *string) (b bool) {
	if IsEmpty(s) {
		return false
	}
	if _, err := strconv.Atoi(*s); err != nil {
		return false
	}
	return true
}

// IsNotInt validates if the string is not a number
func IsNotInt(s *string) (b bool) {
	return !IsInt(s)
}

// ToInt casts the string into a integer
func ToInt(s *string) (v int, err error) {
	return strconv.Atoi(*s)
}

// ToBool casts the string into a bool
func ToBool(s *string) (b bool, err error) {
	return strconv.ParseBool(*s)
}

// ToBoolOrDefault returns the value of the string to bool or the default value in case s has an invalid value
func ToBoolOrDefault(s *string, defaultValue bool) bool {
	val, err := ToBool(s)
	if err != nil {
		return defaultValue
	}
	return val
}
