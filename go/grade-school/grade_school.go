package school

import "sort"

type School struct {
	enroll Grades
}

type Grades []Grade

type Grade struct {
	grade    int
	students []string
}

func (g Grades) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func (g Grades) Len() int {
	return len(g)
}

func (g Grades) Less(i, j int) bool {
	return g[i].grade < g[j].grade
}

func New() *School {
	return &School{}
}

func (s *School) Enrollment() []Grade {
	return s.enroll
}

func (s *School) Add(student string, grade int) {
	for i, g := range s.enroll {
		if g.grade == grade {
			s.enroll[i].students = append(g.students, student)
			sort.Strings(s.enroll[i].students)
			return
		}
	}
	s.enroll = append(s.enroll, Grade{grade, []string{student}})
	sort.Sort(s.enroll)
}

func (s *School) Grade(grade int) []string {
	for _, g := range s.enroll {
		if g.grade == grade {
			return g.students
		}
	}
	return []string{}
}
