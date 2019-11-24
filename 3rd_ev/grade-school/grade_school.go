package school

import "sort"

//Grade
type Grade struct {
	grade int
	names []string
}

//School
type School struct {
	grades map[int]Grade
}

//New
func New() *School {
	return &School{grades: make(map[int]Grade)}
}

//Add
func (s *School) Add(name string, grade int) {
	g := s.grades[grade]
	g.grade, g.names = grade, append(g.names, name)
	s.grades[grade] = g
}

//Grade
func (s *School) Grade(in int) []string {
	return s.grades[in].names
}

//Enrollment
func (s *School) Enrollment() []Grade {
	grades := []Grade{}
	for _, g := range s.grades {
		sort.Strings(g.names)
		if len(grades) == 0 {
			grades = append(grades, g)
			continue
		}
		for i, gg := range grades {
			if g.grade < gg.grade {
				grades = append(grades, Grade{})
				copy(grades[i+1:], grades[i:])
				grades[i] = g
				break
			}
			if i == len(grades)-1 {
				grades = append(grades, g)
			}
		}
	}
	return grades
}