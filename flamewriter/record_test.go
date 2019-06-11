package flamewriter

import "testing"

func TestRecord_ValueStatisticInplace(t *testing.T) {
	var m map[int]int
	m = make(map[int]int, 0)
	makeRecord().ValueStatisticInplace(m)
	if len(m) != 3 {
		t.Errorf("3!=%d", len(m))
		return
	}
	m = make(map[int]int, 0)
	makeRecord().CutoffInplace(3).ValueStatisticInplace(m)
	if len(m) != 3 {
		t.Errorf("3!=%d", len(m))
		return
	}
	m = make(map[int]int, 0)
	makeRecord().CutoffInplace(6).ValueStatisticInplace(m)
	if len(m) != 1 {
		t.Errorf("1!=%d", len(m))
		return
	}
}
