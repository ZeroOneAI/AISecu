package utils

import "errors"

func StringToBool(str string) (bool, error) {
	if str == "TRUE" || str == "True" || str == "true" {
		return true, nil
	}
	if str == "FALSE" || str == "False" || str == "false" {
		return false, nil
	}
	return false, errors.New("Invalid Value [" + str + "]")
}
