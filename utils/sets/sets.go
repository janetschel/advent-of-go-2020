package sets

import (
	"fmt"
	"go/types"
)

type Set struct {
	elements map[string]types.Nil
}

func New() Set {
	set := new(Set)
	set.elements = make(map[string]types.Nil)
	return *set
}

func (set *Set) Add(element string) {
	if element != "" {
		set.elements[element] = types.Nil{}
	}
}

func (set *Set) Remove(element string) {
	if _, isInMap := set.elements[element]; isInMap {
		delete(set.elements, element)
	}
}

func (set *Set) Wipe() {
	set.elements = make(map[string]types.Nil)
}

func (set *Set) Copy() Set {
	resultSet := New()

	for element := range set.elements {
		resultSet.Add(element)
	}

	return resultSet
}

func (set *Set) Has(element string) bool {
	_, has := set.elements[element]
	return has
}

func (set *Set) Size() int {
	return len(set.elements)
}

func (set *Set) Max() string {
	max := ""

	for element := range set.elements {
		if element > max {
			max = element
		}
	}

	return max
}

func (set *Set) Min() string {
	min := ""

	for element := range set.elements {
		if element < min {
			min = element
		}
	}

	return min
}

func (set *Set) Intersect(setToIntersectWith Set) Set {
	resultSet := New()

	for element := range set.elements {
		if set.Has(element) != setToIntersectWith.Has(element) {
			resultSet.Add(element)
		}
	}

	return resultSet
}

func (set Set) String() string {
	str := "{ "
	for element := range set.elements {
		str += fmt.Sprintf("'%s', ", element)
	}

	return str[:len(str)-2] + " }"
}
