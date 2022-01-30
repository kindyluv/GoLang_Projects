package main

type Student struct {
	id           string
	firstName    string
	lastName     string
	age          int
	grades       []int
	cohortName   string
	cohortNumber int
}

func (s *Student) setId(id string) {
	s.id = id
}

func (s *Student) getId() string {
	return s.id
}

func (s *Student) setFirstName(firstName string) {
	s.firstName = firstName
}

func (s *Student) getFirstName() string {
	return s.firstName
}

func (s *Student) setLastName(lastName string) {
	s.lastName = lastName
}

func (s *Student) getLastName() string {
	return s.lastName
}

func (s *Student) setAndGetAge(age int) int {
	s.age = age
	return s.age
}

func (s *Student) getGrade(grade []int) float64 {
	sum := 0
	for _, v := range grade {
		sum += v
	}
	return float64(sum)
}

func (s *Student) setCohortName(cohort string) {
	s.cohortName = cohort
}

func (s *Student) getCohortName() string {
	return s.cohortName
}

func (s *Student) setCohortNumber(cohortNo int) {
	s.cohortNumber = cohortNo
}

func (s *Student) getCohortNumber() int {
	return s.cohortNumber
}
