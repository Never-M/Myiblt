package myiblt

import (
	"errors"
	"fmt"
)

type Cell struct {
	Count   int
	Key_sum int
	Val_sum int
}

type Table struct {
	Cell_len int
	Hash_len int
	Arr      []Cell
}

func Create(cellLen int, hashLen int) *Table {
	table := new(Table)
	table.Cell_len = cellLen
	table.Hash_len = hashLen
	table.Arr = make([]Cell, cellLen)
	return table
}

func (t *Table) hash(id, key, seed int) int {
	return (key*seed + id) % t.Cell_len
}

func (t *Table) bloom(key int) []int {
	idxs := make([]int, t.Hash_len)
	exist := make(map[int]int)

	for i := 0; i < t.Hash_len; i++ {
		var idx int
		id := 0
		for {
			idx = t.hash(id, key, i)
			if _, ok := exist[idx]; ok {
				id++
			} else {
				break
			}
		}
		idxs[i] = idx
		exist[idx] = 1
	}
	return idxs
}

func (t *Table) insert(key, val int) error {
	idxs := t.bloom(key)

	for i := 0; i < len(idxs); i++ {
		t.Arr[idxs[i]].Count++
		t.Arr[idxs[i]].Key_sum += key
		t.Arr[idxs[i]].Val_sum += val
	}
	return nil
}

func (t *Table) delete(key, val int) error {
	idxs := t.bloom(key)

	for i := 0; i < len(idxs); i++ {
		t.Arr[idxs[i]].Count--
		t.Arr[idxs[i]].Key_sum -= key
		t.Arr[idxs[i]].Val_sum -= val
	}
	return nil
}

func (t *Table) get(key int) (int, error) {
	idxs := t.bloom(key)
	for i := 0; i < len(idxs); i++ {
		if t.Arr[idxs[i]].Count == 0 {
			return 0, errors.New("The key " + string(key) + "doesn't exist!")
		} else if t.Arr[idxs[i]].Count == 1 {
			if t.Arr[idxs[i]].Key_sum == key {
				return t.Arr[idxs[i]].Val_sum, nil
			} else {
				return 0, errors.New("The key " + string(key) + "doesn't exist!")
			}
		}
	}
	return 0, errors.New("Not sure whether the key (" + string(key) + ") exist or not.")
}

func (t *Table) listentry() ([][]int, error) {
	var res [][]int
	count0 := 0
	count1 := 0
	for {
		count0 = 0
		count1 = 0
		for i := 0; i < t.Cell_len; i++ {
			if t.Arr[i].Count == 1 {
				count1 = 1
				res = append(res, []int{t.Arr[i].Key_sum, t.Arr[i].Val_sum})
				t.delete(t.Arr[i].Key_sum, t.Arr[i].Val_sum)
			} else if t.Arr[i].Count == 0 {
				count0++
			}
		}
		if count1 == 0 {
			break
		}
	}
	fmt.Println(count0)
	if count0 == t.Cell_len {
		return res, nil
	}
	return nil, errors.New("can't extract from the table")
}
