package change

import "errors"

func Change(coins []int, target int) ([]int, error) {
	result := []int{}

	if target == 0 {
		return result, nil
	}

	if target < 0 {
		return result, errors.New("target amount cannot be negative")
	}

	if target < coins[0] {
		return result, errors.New("target amount is less than the smallest coin value")
	}

	for _, coin := range coins {
		if coin == target {
			return []int{target}, nil
		}
	}

	for count := 0; count < 2; count++ {
		res, s, i := []int{}, target, -1

		for index, coin := range coins {
			if coin > target {
				i = index - 1
				break
			}
		}

		if i < 0 {
			i = len(coins) - 1
		}

		i -= count

		for i >= 0 && s >= coins[0] {
			if s >= coins[i] {
				res = append([]int{coins[i]}, res...)
				s -= res[0]
			}

			if coins[i] > s {
				i -= 1
			} else if i != 0 && (s%coins[i-1] == 0) && (s%coins[i] != 0) {
				i -= 1
			} else if i != 0 && (s > coins[i]) && (s-coins[i] < coins[0]) {
				i -= 1
			}
		}

		if s == 0 && (len(result) == 0 || len(res) < len(result)) {
			result = res
		}
	}

	if len(result) == 0 {
		return result, errors.New("no valid change combination found")
	}

	return result, nil
}
