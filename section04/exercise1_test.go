package section04

import (
	"reflect"
	"sort"
	"testing"
)

func TestNewPerson(t *testing.T) {
	input := NewPerson("Jhon", "Oliver")
	var want *Person
	want = &Person{1, "Jhon", "Oliver"}

	// check the result
	if !reflect.DeepEqual(input, want) {
		t.Error(`NewPerson("Jhon", "Oliver") = {1, "Jhon","Oliver"}`)
	}
}

func TestPersonSlice_Sort(t *testing.T) {
	if personList.Len() != 0 {
		t.Error(`personList.Len() != 0`)
	}

	// input values
	personList = append(personList, NewPerson("Jhon", "Oliver"))
	personList = append(personList, NewPerson("Anne", "Mc"))
	personList = append(personList, NewPerson("Anne", "Anne"))
	personList = append(personList, NewPerson("Anne", "Anne"))

	// want values
	want := PersonSlice{
		{3, "Anne", "Anne"},
		{4, "Anne", "Anne"},
		{2, "Anne", "Mc"},
		{1, "Jhon", "Oliver"},
	}

	// Sort the struct
	sort.Sort(personList)

	// check sort result
	if !reflect.DeepEqual(personList, want) {
		t.Error(`sort.Sort([{1, "Jhon","Oliver"},{2, "Anne","Mc"},{3, "Anne","Anne"},{4, "Anne","Anne"}]) = [{3, "Anne","Anne"},{4, "Anne","Anne"},{2, "Anne","Mc"},{1, "Jhon","Oliver"}]`)
	}
}
