package bookstore

func Cost(basket []int) int {
	var groups, sets [6]int

	for i := range basket {
		groups[basket[i]]++
	}

	for setSize := 0; sets[0] == 0; sets[setSize], setSize = sets[setSize]+1, 0 {
		for i := range groups {
			if groups[i] > 0 {
				setSize++
				groups[i]--
			}
		}
	}

	for sets[3] > 0 && sets[5] > 0 {
		sets[3], sets[5], sets[4] = sets[3]-1, sets[5]-1, sets[4]+2
	}

	return 800*sets[1] + (800*2*0.95)*sets[2] + (800*3*0.9)*sets[3] + (800*4*0.80)*sets[4] + (800*5*0.75)*sets[5]
}
