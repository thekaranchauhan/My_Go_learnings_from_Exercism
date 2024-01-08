package stringset

import "fmt"

type Set map[string]bool

func New() Set { return make(map[string]bool) }

func NewFromSlice(slice []string) Set {
	set := New()
	if slice == nil {
		return set
	}
	for _, s := range slice {
		set[s] = true
	}
	return set
}

func (set Set) String() string {
	s := ""
	for k, _ := range set {
		if len(s) > 0 {
			s += ", "
		}
		s += fmt.Sprintf("\"%s\"", k)
	}
	return fmt.Sprintf("{%s}", s)
}

func (set Set) IsEmpty() bool     { return len(set) == 0 }
func (set Set) Has(s string) bool { return set[s] }
func (set Set) Add(s string)      { set[s] = true }

func Subset(s1, s2 Set) bool {
	for k, _ := range s1 {
		if !s2[k] {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for k, _ := range s1 {
		if s2[k] {
			return false
		}
	}
	for k, _ := range s2 {
		if s1[k] {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool { return Subset(s1, s2) && Subset(s2, s1) }

func Intersection(s1, s2 Set) Set {
	inter := New()
	for k, _ := range s1 {
		if s2[k] {
			inter[k] = true
		}
	}
	return inter
}

func Difference(s1, s2 Set) Set {
	diff := New()
	for k, _ := range s1 {
		if !s2[k] {
			diff[k] = true
		}
	}
	return diff
}

func Union(s1, s2 Set) Set {
	union := New()
	for k, _ := range s1 {
		union[k] = true
	}
	for k, _ := range s2 {
		union[k] = true
	}
	return union
}
