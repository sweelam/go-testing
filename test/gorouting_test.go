package test

import "testing"

func TestGoRoutine(t *testing.T) {
	ch := make(chan uint64)

	list := []int{1, 2, 3, 4, 5}
	go sumList(ch, list)

	sum := <-ch
	if sum != 15 {
		t.Errorf("Expected 15, got %d", sum)
	}
}

func TestGoRoutineLoader(t *testing.T) {
	chSize := 1000_000_00
	ch := make([]chan uint64, chSize)

	list := []int{1, 2, 3, 4, 5}
	for i := range chSize {
		ch[i] = make(chan uint64, 1)
		go sumList(ch[i], list)
	}

	if len(ch) != chSize {
		t.Errorf("Expected chSize, got %d", len(ch))
	}

	for i := range chSize {
		sum := <-ch[i]
		if sum != 15 {
			t.Errorf("Expected 15, got %d", sum)
		}
	}
}

func sumList(ch chan uint64, list []int) {
	var sum uint64

	for _, s := range list {
		sum += uint64(s)
	}

	ch <- sum
}
