package iossimulator

import "regexp"

func ParseRuntime(str string) (string, error) {
	runtimeRegex, compileErr := regexp.Compile(`(iOS (\d+.\d)|watchOS (\d+.\d)|tvOS (\d+.\d))`)
	if compileErr != nil {
		return "", compileErr
	}

	return runtimeRegex.FindString(str), nil
}

func ParseRuntimeId(str string) (string, error) {
	runtimeIdRegex, compileErr := regexp.Compile(`(com\.apple\.[^\s\(\)]+)`)
	if compileErr != nil {
		return "", compileErr
	}

	return runtimeIdRegex.FindString(str), nil
}
