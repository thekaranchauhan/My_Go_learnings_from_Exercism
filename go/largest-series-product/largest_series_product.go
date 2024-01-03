package lsproduct

import (
	"errors"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if len(digits) < span || span < 0 {
		return 0, errors.New("invalid span")
	}
	var digitList []int64
	for _, v := range digits {
		if v < '0' || v > '9' {
			return 0, errors.New("invalid digits")
		}
		digitList = append(digitList, int64(v-'0'))
	}
	var largest int64
	for i := 0; i+span <= len(digitList); i++ {
		tmp := int64(1)
		for _, v := range digitList[i : i+span] {
			tmp *= v
		}
		if tmp > largest {
			largest = tmp
		}
	}
	return largest, nil
}
