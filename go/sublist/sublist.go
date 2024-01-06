package sublist

import (
	"strconv"
	"strings"
)

// Relation type is defined in relations.go file.
func Sublist(l1, l2 []int) Relation {
	if len(l1) == 0 && len(l2) == 0 {
		return RelationEqual
	}
	if len(l1) > 0 && len(l2) == 0 {
		return RelationSuperlist
	}
	if len(l1) == 0 && len(l2) > 0 {
		return RelationSublist
	}
	s1, s2 := "", ""
	for _, num := range l1 {
		s1 += strconv.Itoa(num) + ","
	}
	for _, num := range l2 {
		s2 += strconv.Itoa(num) + ","
	}
	if s1 == s2 {
		return RelationEqual
	}
	if strings.Contains(s1, s2) {
		return RelationSuperlist
	}
	if strings.Contains(s2, s1) {
		return RelationSublist
	}
	return RelationUnequal
}
