package myiblt

import (
	"testing"
)

func Test_Insert(t *testing.T) {
	table := Create(7, 4)
	table.insert(2, 5)
	table.insert(7, 9)
	val, err := table.get(2)
	assertEqual(t, val, 5)
	val, err = table.get(7)
	assertEqual(t, val, 9)
	table.delete(2, 5)
	val, err = table.get(2)
	assertNotNil(t, err)
}
