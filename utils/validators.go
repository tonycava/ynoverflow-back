package utils

import "github.com/asaskevich/govalidator"

func IsEmpty(str string) (bool, string) {
	if govalidator.HasWhitespaceOnly(str) && len(str) == 0 {
		return true, "cannot be empty"
	}
	return false, ""
}
