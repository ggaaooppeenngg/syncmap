package syncmap

import (
	"testing"
)

func TestMap(t *testing.T) {
	m := New[int, string](map[int]string{1: "1"})
	v, ok := m.Get(2)
	if ok {
		t.Fatalf("get %s\n", v)
	}
	m.Set(2, "2")
	v, ok = m.Get(2)
	if ok && v != "2" {
		t.Fatalf("get %s\n", v)
	}
	dupMap := m.Dup()
	v, ok = dupMap.Get(1)
	if ok && v != "1" {
		t.Fatalf("get %s\n", v)
	}
	m.Update(3, "3")
	v, ok = m.Get(3)
	if ok && v != "3" {
		t.Fatalf("get %s\n", v)
	}
	ok = m.Update(4, "4")
	if ok {
		t.Fatalf("get %t\n", ok)
	}
	m.Delete(3)
	v, ok = m.Get(3)
	if ok && v != "" {
		t.Fatalf("get %s\n", v)
	}
}
