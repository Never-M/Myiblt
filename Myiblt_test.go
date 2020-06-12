package myiblt

import (
	"testing"
)

func TestInsertDeleteGet(t *testing.T) {
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

func TestListEntry(t *testing.T) {
	AliceTable := Create(7, 4)
	AliceTable.insert(2, 5)
	AliceTable.insert(7, 9)
	BobTable := Create(7, 4)
	BobTable.insert(2, 5)
	BobTable.insert(7, 9)

	BobArr, _ := BobTable.listentry()

	for _, pair := range BobArr {
		AliceTable.delete(pair[0], pair[1])
	}
	AliceArr, _ := AliceTable.listentry()
	assertEqual(t, len(AliceArr), 0)
}
