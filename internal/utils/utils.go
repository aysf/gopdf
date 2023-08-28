package utils

import (
	"errors"
	"path/filepath"
	"strconv"
	"strings"
)

func FileNameWithoutExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func Flatten2D(slice [][]int) ([]int, error) {
	var result []int

	for _, innerSlice := range slice {
		if len(innerSlice) == 2 {
			start, end := innerSlice[0], innerSlice[1]

			if start > end {
				return nil, errors.New("invalid input: start value is greater than end value")
			}

			for i := start; i <= end; i++ {
				result = append(result, i)
			}
		} else if len(innerSlice) == 1 {
			result = append(result, innerSlice[0])
		} else {
			return nil, errors.New("invalid input: inner slice must have 1 or 2 elements")
		}
	}

	// Check for duplicate elements
	seen := make(map[int]bool)
	for _, num := range result {
		if seen[num] {
			return nil, errors.New("invalid input: non-unique element in the result slice")
		}
		seen[num] = true
	}

	return result, nil
}

func StrTo2DStr(input string) ([][]string, error) {
	groups := strings.Split(input, ",")
	result := make([][]string, len(groups))

	for i, group := range groups {
		ranges := strings.Split(group, "-")
		start := ranges[0]

		if len(ranges) == 1 {
			result[i] = []string{start}
		} else {
			end := ranges[1]
			result[i] = []string{start, end}
		}
	}
	return result, nil
}

func StrToInt2D(str2D [][]string) ([][]int, error) {
	return nil, nil
}

func Flatten2DStrToInt(strings [][]string) ([]int, error) {
	var result []int

	for _, innerSlice := range strings {
		if len(innerSlice) == 2 {
			startStr, endStr := innerSlice[0], innerSlice[1]

			start, err := strconv.Atoi(startStr)
			if err != nil {
				return nil, err
			}
			end, err := strconv.Atoi(endStr)
			if err != nil {
				return nil, err
			}

			if start > end {
				return nil, errors.New("invalid input: start value is greater than end value")
			}

			for i := start; i <= end; i++ {
				result = append(result, i)
			}
		} else if len(innerSlice) == 1 {
			num, err := strconv.Atoi(innerSlice[0])
			if err != nil {
				return nil, err
			}
			result = append(result, num)
		} else {
			return nil, errors.New("invalid input: inner slice must have 1 or 2 elements")
		}
	}

	// Check for duplicate elements
	seen := make(map[int]bool)
	for _, num := range result {
		if seen[num] {
			return nil, errors.New("invalid input: non-unique element in the result slice")
		}
		seen[num] = true
	}

	return result, nil
}
