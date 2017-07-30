package section04

// Problem 1: Sorting Names
// Sorting in Go is done through interfaces!
// To sort a collection (such as a slice), the type must satisfy sort.Interface,
// which requires 3 methods: Len() int, Less(i, j int) bool, and Swap(i, j int).
// To actually sort a slice, you need to first implement all 3 methods on a
// custom type, and then call sort.Sort on your the PersonSlice type.
// See the Go documentation: https://golang.org/pkg/sort/ for full details.

// Person stores a simple profile. These should be sorted by alphabetical order
// by last name, followed by the first name, followed by the ID. You can assume
// the ID will be unique, but the names need not be unique.
// Sorting should be case-sensitive and UTF-8 aware.
type Person struct {
	ID        int
	FirstName string
	LastName  string
}

// PersonSlice is a list of *Person
type PersonSlice []*Person

// Len is part of sort.Interface.
func (a PersonSlice) Len() int { return len(a) }

// Swap is part of sort.Interface.
func (a PersonSlice) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less is part of sort.Interface.
func (a PersonSlice) Less(i, j int) bool {
	if a[i].LastName < a[j].LastName {
		return true
	}
	if a[i].LastName > a[j].LastName {
		return false
	}
	return a[i].FirstName < a[j].FirstName
}

var personList PersonSlice

// NewPerson is a constructor for Person. ID should be assigned automatically in
// sequential order, starting at 1 for the first Person created.
func NewPerson(first, last string) *Person {
	var p Person
	p.ID = personList.Len() + 1
	p.FirstName = first
	p.LastName = last
	return &p
}
